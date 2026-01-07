package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/app/model/request"
)

type AgentTransactionService struct{}

// GetAgentTransactionList 分页获取代理分润记录列表
func (s *AgentTransactionService) GetAgentTransactionList(info request.AgentTransactionSearch) (list []model.AgentTransaction, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.AgentTransaction{})

	// 查询条件
	if info.MerchantId != nil {
		db = db.Where("merchant_id = ?", info.MerchantId)
	}
	if info.OrderNo != nil && *info.OrderNo != "" {
		db = db.Where("order_no LIKE ?", "%"+*info.OrderNo+"%")
	}
	if info.BeneficiaryId != nil {
		db = db.Where("beneficiary_id = ?", info.BeneficiaryId)
	}
	if info.SourceUserId != nil {
		db = db.Where("source_user_id = ?", info.SourceUserId)
	}
	if info.AgentLevelName != nil && *info.AgentLevelName != "" {
		db = db.Where("agent_level_name LIKE ?", "%"+*info.AgentLevelName+"%")
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("created_at desc").Find(&list).Error
	return
}

// ProcessOrderPayment 处理订单支付后的代理逻辑
func (s *AgentTransactionService) ProcessOrderPayment(ctx context.Context, order model.Order) error {
	if order.MerchantID == nil || order.UserID == nil {
		return nil
	}
	if order.PayStatus == nil || *order.PayStatus != "paid" {
		return nil
	}

	// 1. 检查并升级代理等级 (从下往上触发，影响整条链)
	// 需要检查下单用户及其所有上级
	if err := s.CheckAndUpgradeAgents(ctx, order); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("CheckAndUpgradeAgents error: %v", err))
	}

	// 2. 分发分润
	if err := s.DistributeCommissions(ctx, order); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("DistributeCommissions error: %v", err))
	}

	return nil
}

// CheckAndUpgradeAgents 检查并升级代理等级
func (s *AgentTransactionService) CheckAndUpgradeAgents(ctx context.Context, order model.Order) error {
	merchantID := uint(*order.MerchantID)
	userID := uint(*order.UserID)

	// 获取用户信息以拿到邀请路径
	var user model.AppUsers
	if err := global.GVA_DB.WithContext(ctx).First(&user, userID).Error; err != nil {
		return err
	}

	// 目标用户列表：自己 + 所有上级
	targetUserIDs := []uint{userID}
	if user.InvitePath != nil && *user.InvitePath != "" {
		ids := strings.Split(*user.InvitePath, "/")
		for _, idStr := range ids {
			if idStr == "" {
				continue
			}
			if uid, err := strconv.ParseUint(idStr, 10, 64); err == nil {
				targetUserIDs = append(targetUserIDs, uint(uid))
			}
		}
	}

	// 获取该商户的所有代理等级配置，按金额降序排列
	var levels []model.AgentLevel
	if err := global.GVA_DB.WithContext(ctx).Where("merchant_id = ? AND status = '1'", merchantID).Order("qualification_amount desc").Find(&levels).Error; err != nil {
		return err
	}
	if len(levels) == 0 {
		return nil
	}

	// 遍历每个用户进行检查
	for _, targetUID := range targetUserIDs {
		// 1. 计算总业绩 (自己 + 所有下级)
		performance, err := s.CalculateTeamPerformance(ctx, targetUID, merchantID)
		if err != nil {
			continue
		}

		// 2. 匹配最高等级
		// performance 是 float64, QualificationAmount 是 int32 (万)
		performanceWan := performance / 10000.0
		var matchedLevel *model.AgentLevel

		for i := range levels {
			if levels[i].QualificationAmount != nil && performanceWan >= float64(*levels[i].QualificationAmount) {
				matchedLevel = &levels[i]
				break // 已经是降序，找到第一个即为最高
			}
		}

		if matchedLevel == nil {
			continue
		}

		// 3. 检查是否需要更新/插入
		var currentUAL model.UserAgentLevel
		err = global.GVA_DB.WithContext(ctx).Where("user_id = ? AND merchant_id = ?", targetUID, merchantID).First(&currentUAL).Error
		
		shouldSave := false
		if err != nil {
			// 不存在，创建
			currentUAL = model.UserAgentLevel{
				UserID:         targetUID,
				MerchantID:     merchantID,
				AgentLevelID:   uint(*matchedLevel.Id),
				AgentLevelName: *matchedLevel.Name,
				Status:         "active",
				AchievedAt:     time.Now(),
			}
			shouldSave = true
		} else {
			// 已存在，检查是否升级
			// 简单起见，比较 QualificationAmount
			var currentLevelConfig model.AgentLevel
			if err := global.GVA_DB.WithContext(ctx).First(&currentLevelConfig, currentUAL.AgentLevelID).Error; err == nil {
				if currentLevelConfig.QualificationAmount != nil && *matchedLevel.QualificationAmount > *currentLevelConfig.QualificationAmount {
					currentUAL.AgentLevelID = uint(*matchedLevel.Id)
					currentUAL.AgentLevelName = *matchedLevel.Name
					currentUAL.AchievedAt = time.Now()
					shouldSave = true
				}
			} else {
				// 如果找不到旧配置，直接更新
				currentUAL.AgentLevelID = uint(*matchedLevel.Id)
				currentUAL.AgentLevelName = *matchedLevel.Name
				shouldSave = true
			}
		}

		if shouldSave {
			if currentUAL.ID > 0 {
				global.GVA_DB.WithContext(ctx).Save(&currentUAL)
			} else {
				global.GVA_DB.WithContext(ctx).Create(&currentUAL)
			}
		}
	}

	return nil
}

// CalculateTeamPerformance 计算团队总业绩 (自己 + 下级)
func (s *AgentTransactionService) CalculateTeamPerformance(ctx context.Context, userID uint, merchantID uint) (float64, error) {
	// 查找所有下级ID
	// invite_path LIKE 'uid/%' OR invite_path LIKE '%/uid/%' OR invite_path LIKE '%/uid'
	uidStr := strconv.Itoa(int(userID))
	likeStart := uidStr + "/%"
	likeMid := "%/" + uidStr + "/%"
	likeEnd := "%/" + uidStr

	var subUserIDs []uint
	// 包含自己
	subUserIDs = append(subUserIDs, userID)

	var subs []model.AppUsers
	err := global.GVA_DB.WithContext(ctx).Select("id").Where("merchant_id = ? AND (invite_path LIKE ? OR invite_path LIKE ? OR invite_path LIKE ?)", merchantID, likeStart, likeMid, likeEnd).Find(&subs).Error
	if err != nil {
		return 0, err
	}
	for _, sub := range subs {
		subUserIDs = append(subUserIDs, sub.ID)
	}

	var total float64
	// 统计所有有效订单
	// status IN ('paid', 'shipped', 'completed')
	err = global.GVA_DB.WithContext(ctx).Model(&model.Order{}).
		Where("merchant_id = ? AND user_id IN ? AND pay_status = 'paid'", merchantID, subUserIDs).
		Select("COALESCE(SUM(total_amount), 0)").
		Scan(&total).Error

	return total, err
}

// DistributeCommissions 分发分润
func (s *AgentTransactionService) DistributeCommissions(ctx context.Context, order model.Order) error {
	merchantID := uint(*order.MerchantID)
	sourceUserID := uint(*order.UserID)

	// 基础金额 87%
	orderTotal := 0.0
	if order.TotalAmount != nil {
		orderTotal = *order.TotalAmount
	}
	baseAmount := orderTotal * 0.87

	// 获取用户及路径
	var user model.AppUsers
	if err := global.GVA_DB.WithContext(ctx).First(&user, sourceUserID).Error; err != nil {
		return err
	}

	// 路径：Source -> Parent -> GrandParent ...
	// 需要反转 invite_path (它是 1/2/3 形式，3是parent，1是root) -> [1, 2, 3]
	// 实际上，InvitePath 存储的是 Root -> Parent。
	// 比如 Root(1) -> A(2) -> B(3) -> Source(4). Source.InvitePath = "1/2/3".
	// 我们需要从 3 -> 2 -> 1 的顺序遍历。
	// 另外，SourceUser 自己也可能是代理，需要参与 'allSubordinates' 的判断吗？
	// 需求 4.4.2: "A4 is AGENT1, A3 is AGENT2". A4 is the buyer.
	// 所以需要把 SourceUser 也加入到遍历列表的最前面。

	var candidateIDs []uint
	candidateIDs = append(candidateIDs, sourceUserID) // 先加入自己

	if user.InvitePath != nil && *user.InvitePath != "" {
		ids := strings.Split(*user.InvitePath, "/")
		// 倒序加入
		for i := len(ids) - 1; i >= 0; i-- {
			if ids[i] == "" {
				continue
			}
			if uid, err := strconv.ParseUint(ids[i], 10, 64); err == nil {
				candidateIDs = append(candidateIDs, uint(uid))
			}
		}
	}

	// 记录上一次分润的等级权重 (QualificationAmount)
	// 用于 'allSubordinates' 的级差逻辑
	var lastPaidWeight int32 = -1

	for _, uid := range candidateIDs {
		// 1. 获取该用户的代理等级
		var ual model.UserAgentLevel
		if err := global.GVA_DB.WithContext(ctx).Where("user_id = ? AND merchant_id = ? AND status = 'active'", uid, merchantID).First(&ual).Error; err != nil {
			// 用户不是代理，跳过
			continue
		}

		// 2. 获取等级配置
		var levelConfig model.AgentLevel
		if err := global.GVA_DB.WithContext(ctx).First(&levelConfig, ual.AgentLevelID).Error; err != nil {
			continue
		}

		scope := "directly"
		if levelConfig.DividendScope != nil {
			scope = *levelConfig.DividendScope
		}

		shouldPay := false

		// 3. 判断是否分润
		if scope == "directly" {
			// 只有直属上级分润
			// user.InviterID 必须等于 uid
			// 注意：candidateIDs[0] 是 sourceUser 自己。自己不能是自己的直属上级。
			if user.InviterID != nil && uint(*user.InviterID) == uid {
				shouldPay = true
			}
		} else if scope == "allSubordinates" || scope == "all" { // 兼容
			// 级差逻辑
			currentWeight := int32(0)
			if levelConfig.QualificationAmount != nil {
				currentWeight = *levelConfig.QualificationAmount
			}

			// 如果当前等级 > 上次分润等级 (或者还没分过)
			// 注意：i==0 是 sourceUser 自己。
			// 如果 sourceUser 自己是代理，他能拿自己的提成吗？
			// 需求 4.3.2 例子：A4(Buyer) -> A3. A4 is AGENT1. A3 is AGENT2.
			// A4 gets AGENT1 rate. A3 gets AGENT2 rate.
			// So yes, SourceUser can get commission.
			
			if currentWeight > lastPaidWeight {
				shouldPay = true
				lastPaidWeight = currentWeight
			}
		}

		if shouldPay {
			// 4. 计算金额
			orderNo := ""
			if order.OrderNo != nil {
				orderNo = *order.OrderNo
			}
			tx := model.AgentTransaction{
				OrderNo:        orderNo,
				OrderID:        order.ID,
				OrderAmount:    orderTotal,
				BaseAmount:     baseAmount,
				MerchantID:     merchantID,
				BeneficiaryID:  uid,
				SourceUserID:   sourceUserID,
				AgentLevelID:   ual.AgentLevelID,
				AgentLevelName: ual.AgentLevelName,
				DividendScope:  scope,
			}

			tx.Rate1 = valOrZero(levelConfig.CfRate1)
			tx.Amount1 = baseAmount * tx.Rate1
			
			tx.Rate2 = valOrZero(levelConfig.CfRate2)
			tx.Amount2 = baseAmount * tx.Rate2
			
			tx.Rate3 = valOrZero(levelConfig.CfRate3)
			tx.Amount3 = baseAmount * tx.Rate3
			
			tx.Rate4 = valOrZero(levelConfig.CfRate4)
			tx.Amount4 = baseAmount * tx.Rate4
			
			tx.Rate5 = valOrZero(levelConfig.CfRate5)
			tx.Amount5 = baseAmount * tx.Rate5

			tx.TotalAmount = tx.Amount1 + tx.Amount2 + tx.Amount3 + tx.Amount4 + tx.Amount5
			
			if tx.TotalAmount > 0 {
				global.GVA_DB.WithContext(ctx).Create(&tx)
			}
		}
	}

	return nil
}

func valOrZero(ptr *float64) float64 {
	if ptr == nil {
		return 0
	}
	return *ptr
}
