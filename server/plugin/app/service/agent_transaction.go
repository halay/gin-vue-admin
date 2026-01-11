package service

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"

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

	// 3. 经销商分润 (新增逻辑)
	if err := s.DistributeDealerCommissions(ctx, order); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("DistributeDealerCommissions error: %v", err))
	}

	// 4. 股东分润 (新增逻辑)
	if err := s.DistributeShareholderCommissions(ctx, order); err != nil {
		global.GVA_LOG.Error(fmt.Sprintf("DistributeShareholderCommissions error: %v", err))
	}

	return nil
}

// DistributeShareholderCommissions 分发股东分润
func (s *AgentTransactionService) DistributeShareholderCommissions(ctx context.Context, order model.Order) error {
	// 1. 基础校验
	if order.MerchantID == nil {
		return nil
	}

	merchantID := uint(*order.MerchantID)

	// 2. 计算分润基础金额 A
	// app_orders.total_amount * app_products.tax_rate

	baseAmountA := 0.0
	var items []model.OrderItem
	if err := global.GVA_DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items).Error; err != nil {
		return err
	}

	for _, item := range items {
		// 获取商品税率
		taxRate := 0.0
		var prod model.Product
		if err := global.GVA_DB.WithContext(ctx).Select("tax_rate").First(&prod, item.ProductID).Error; err == nil {
			if prod.TaxRate != nil {
				taxRate = *prod.TaxRate
			}
		}

		itemTotal := 0.0
		if item.TotalAmount != nil {
			itemTotal = *item.TotalAmount
		}

		// 分润基础金额A += 商品总额 * 税率 (百分比)
		baseAmountA += itemTotal * (taxRate / 100.0)
	}

	if baseAmountA <= 0 {
		return nil
	}

	// 3. 获取该商户下所有的股东分润配置
	var shareholders []model.ShareholderProfit
	if err := global.GVA_DB.WithContext(ctx).Where("merchant_id = ?", merchantID).Find(&shareholders).Error; err != nil {
		return err
	}

	if len(shareholders) == 0 {
		return nil
	}

	// 4. 遍历每个股东配置，找到对应的用户并分润
	for _, shareholder := range shareholders {
		shareRatio := 0.0
		if shareholder.ShareRatio != nil {
			shareRatio = *shareholder.ShareRatio
		}

		// 5. 查找绑定了该股东身份的所有用户
		var users []model.AppUsers
		if err := global.GVA_DB.WithContext(ctx).Where("shareholder_profit_id = ?", shareholder.ID).Find(&users).Error; err != nil {
			continue
		}

		userCount := len(users)
		if userCount == 0 {
			continue
		}

		// 6. 计算每个用户分得的金额
		// 总分润 = 基础金额A * share_ratio / 100
		// 每人分润 = 总分润 / userCount

		totalProfit := baseAmountA * (shareRatio / 100.0)
		perUserProfit := totalProfit / float64(userCount)

		if perUserProfit <= 0 {
			continue
		}

		// 7. 为每个用户创建分润记录
		orderNo := ""
		if order.OrderNo != nil {
			orderNo = *order.OrderNo
		}

		for _, user := range users {
			tx := model.AgentTransaction{
				OrderNo:        orderNo,
				OrderID:        order.ID,
				OrderAmount:    *order.TotalAmount,
				BaseAmount:     baseAmountA,
				MerchantID:     merchantID,
				BeneficiaryID:  user.ID,             // 受益人
				SourceUserID:   uint(*order.UserID), // 来源是下单用户
				AgentLevelID:   shareholder.ID,      // 借用字段存 ShareholderID
				AgentLevelName: *shareholder.Name,   // 借用字段存 ShareholderName
				DividendScope:  "shareholder",       // 标记为股东分润
				Description:    fmt.Sprintf("股东分润: 身份[%s]总比例%.2f%%, 共%d人平分", *shareholder.Name, shareRatio, userCount),
				Source:         "shareholder",
				TotalAmount:    perUserProfit,

				// 借用 Rate1/Amount1 存分润信息
				Rate1:   shareRatio,
				Amount1: perUserProfit,
			}

			if err := global.GVA_DB.WithContext(ctx).Create(&tx).Error; err != nil {
				global.GVA_LOG.Error("创建股东分润记录失败", zap.Error(err))
			}
		}
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

	// 基础金额
	orderTotal := 0.0
	if order.TotalAmount != nil {
		orderTotal = *order.TotalAmount
	}
	// 计算税后分润基数: TotalAmount * (1 - TaxRate) * 0.87 (保留原有0.87逻辑吗？用户说改成(1-税点%)，之前是0.87，现在需要替换掉0.87)
	// 用户需求：DistributeCommissions方法里写死的0.87要改成(1-商品的税点*%)
	// 这里需要获取订单对应的商品的税点。订单可能包含多个商品，每个商品税点可能不同。
	// 但 AgentTransaction 是基于 Order 维度的。
	// 简化处理：假设订单中所有商品税点一致，或者取平均？
	// 更准确的做法：遍历 OrderItems，累加每个 Item 的 (ItemAmount * (1 - ItemTaxRate))。

	// 重新计算 BaseAmount
	baseAmount := 0.0
	var items []model.OrderItem
	if err := global.GVA_DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items).Error; err != nil {
		return err
	}

	for _, item := range items {
		itemTotal := 0.0
		if item.TotalAmount != nil {
			itemTotal = *item.TotalAmount
		}

		// 获取商品税点
		taxRate := 0.0
		var prod model.Product
		if err := global.GVA_DB.WithContext(ctx).Select("tax_rate").First(&prod, item.ProductID).Error; err == nil {
			if prod.TaxRate != nil {
				taxRate = *prod.TaxRate
			}
		}

		// 计算该商品的基数: 金额 * (1 - 税率/100)
		baseAmount += itemTotal * (1 - taxRate/100.0)
	}

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

// DistributeDealerCommissions 分发经销商分润
func (s *AgentTransactionService) DistributeDealerCommissions(ctx context.Context, order model.Order) error {
	// 1. 基础校验
	if order.MerchantID == nil || order.UserID == nil {
		return nil
	}

	// 只有已发货(delivered)或已完成(completed)的状态才分润？
	// 需求: "当delivery_status=delivered时触发分润给经销商"
	// 但入口 ProcessOrderPayment 是在 PayStatus=paid 时触发的。
	// 如果需要在 delivered 时触发，则需要修改 ProcessOrderPayment 的调用时机，或者在这里判断状态。
	// 目前 ProcessOrderPayment 是在 PayCallbackCard 和 PayOrderByPoints 中调用的，此时 delivery_status 可能是 none。
	// 所以需要确认：是在支付成功时就分润，还是发货后分润？
	// 需求明确说：apiOrder.UpdateOrder请求时，当delivery_status=delivered时触发
	// 所以我们需要在 UpdateOrder 中埋点，而不是复用 ProcessOrderPayment。
	// 但为了复用代码结构，我们可以独立一个方法，供 UpdateOrder 调用。

	// 这里我们先实现核心逻辑

	merchantID := uint(*order.MerchantID)

	// 2. 计算分润基础金额 A
	// app_orders.total_amount * app_products.tax_rate
	// 由于订单可能有多个商品，我们需要遍历 OrderItems

	baseAmountA := 0.0
	var items []model.OrderItem
	if err := global.GVA_DB.WithContext(ctx).Where("order_id = ?", order.ID).Find(&items).Error; err != nil {
		return err
	}

	for _, item := range items {
		// 获取商品税率
		taxRate := 0.87
		var prod model.Product
		if err := global.GVA_DB.WithContext(ctx).Select("tax_rate").First(&prod, item.ProductID).Error; err == nil {
			if prod.TaxRate != nil {
				taxRate = *prod.TaxRate
			}
		}

		itemTotal := 0.0
		if item.TotalAmount != nil {
			itemTotal = *item.TotalAmount
		}

		// 分润基础金额A += 商品总额 * 税率 (注意税率单位，通常是百分比？需求说 tax_rate，假设是小数或百分比)
		// 如果 tax_rate 存的是 10 表示 10%，则需要 / 100
		// 根据之前的代码 logic: baseAmount += itemTotal * (1 - taxRate/100.0) 推测 taxRate 是 0-100 的值
		// 需求: app_orders.total_amount * app_products.tax_rate
		// 假设 tax_rate 是百分比，例如 10，则 * 0.1
		baseAmountA += itemTotal * (taxRate / 100.0)
	}

	if baseAmountA <= 0 {
		return nil
	}

	// 3. 寻找受益经销商
	// B: 下单用户 (order.UserID)
	// C: B的直属上级 (B.InviterID)

	var userB model.AppUsers
	if err := global.GVA_DB.WithContext(ctx).First(&userB, order.UserID).Error; err != nil {
		return err
	}

	// 如果没有上级，是否需要分润？需求说 "再通过B.inviter_id查询...C"，隐含必须有上级
	if userB.InviterID == nil {
		// 没有上级C，逻辑结束？
		// 需求 4: "再通过B.inviter_id查询...C...如果C.app_users.app_dealer_id为空，则查询app_dealers表的此订单商家的默认配置"
		// 这句话的前提是有 C。如果连 C 都没有，是否直接找默认配置？
		// 根据语境，应该是必须通过 C 来找经销商。如果没有 C，就没有 C.app_dealer_id。
		// 但为了健壮性，如果 B 没有邀请人，也许可以直接跳到默认配置？
		// 严格按照需求：必须有 C。
		return nil
	}

	var userC model.AppUsers
	if err := global.GVA_DB.WithContext(ctx).First(&userC, userB.InviterID).Error; err != nil {
		return err
	}

	var dealer model.AppDealer
	foundDealer := false

	// 4.1 检查 C 的经销商配置
	if userC.AppDealerID != nil {
		if err := global.GVA_DB.WithContext(ctx).First(&dealer, userC.AppDealerID).Error; err == nil {
			foundDealer = true
		}
	}

	// 4.2 如果没找到，使用默认配置
	if !foundDealer {
		if err := global.GVA_DB.WithContext(ctx).Where("merchant_id = ? AND is_default = ?", merchantID, true).First(&dealer).Error; err == nil {
			foundDealer = true
		}
	}

	if !foundDealer {
		return nil
	}

	// 5. 计算分润金额
	// commission_type: 1:比例, 2:固定金额
	// allowance_type: 1:比例, 2:固定金额

	commissionAmount := 0.0
	allowanceAmount := 0.0

	// 销售提成
	salesComm := 0.0
	if dealer.SalesCommission != nil {
		salesComm = *dealer.SalesCommission
	}
	if dealer.CommissionType != nil && *dealer.CommissionType == "2" {
		// 固定金额
		commissionAmount = salesComm
	} else {
		// 比例 (默认)
		commissionAmount = baseAmountA * (salesComm / 100.0)
	}

	// 费用补贴
	expenseAllow := 0.0
	if dealer.ExpenseAllowance != nil {
		expenseAllow = *dealer.ExpenseAllowance
	}
	if dealer.AllowanceType != nil && *dealer.AllowanceType == "2" {
		// 固定金额
		allowanceAmount = expenseAllow
	} else {
		// 比例 (默认)
		allowanceAmount = baseAmountA * (expenseAllow / 100.0)
	}

	totalDealerAmount := commissionAmount + allowanceAmount
	if totalDealerAmount <= 0 {
		return nil
	}

	// 6. 保存记录
	// 经销商分润记录到 app_agent_transactions 表中
	// 使用 Source = 'dealer' 区分？或者 Description

	orderNo := ""
	if order.OrderNo != nil {
		orderNo = *order.OrderNo
	}

	// 经销商关联的用户是谁？需求没明确说受益人ID填谁。
	// 通常经销商是关联到某个账户的，或者直接关联到 C？
	// 需求 4: "C.app_users.app_dealer_id...查询app_dealers表信息"
	// AppDealer 表没有直接绑定 UserID，只有 MerchantID。
	// 但 C 是 "直属上级用户"。
	// 这里的受益人应该是 C 吗？
	// 逻辑是：C 作为上级，根据他绑定的 Dealer 配置（或默认配置）获得分润。
	// 所以 BeneficiaryID = C.ID

	tx := model.AgentTransaction{
		OrderNo:        orderNo,
		OrderID:        order.ID,
		OrderAmount:    *order.TotalAmount,
		BaseAmount:     baseAmountA,
		MerchantID:     merchantID,
		BeneficiaryID:  userC.ID,            // 受益人是 C
		SourceUserID:   uint(*order.UserID), // 来源是 B
		AgentLevelID:   dealer.ID,           // 借用字段存 DealerID
		AgentLevelName: *dealer.Name,        // 借用字段存 DealerName
		DividendScope:  "dealer",            // 标记为经销商分润
		Description:    fmt.Sprintf("经销商分润: 销售提成=%.2f, 费用补贴=%.2f", commissionAmount, allowanceAmount),
		Source:         "dealer",
		TotalAmount:    totalDealerAmount,

		// 借用 Rate1/Amount1 存销售提成
		Rate1:   salesComm,
		Amount1: commissionAmount,

		// 借用 Rate2/Amount2 存费用补贴
		Rate2:   expenseAllow,
		Amount2: allowanceAmount,
	}

	if err := global.GVA_DB.WithContext(ctx).Create(&tx).Error; err != nil {
		return err
	}

	return nil
}
