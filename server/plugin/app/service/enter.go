package service

var Service = new(service)

type service struct {
	AppUsers               appUsers
	Merchants              mc
	Banner                 banner
	MerchantAnnouncement   MA
	MerchantAdmin          MADM
	ProductCategory        PC
	Product                P
	ProductSku             SKU
	ProductSpec            PS
	ProductSpecOption      PSO
	Order                  ORD
	OrderItem              ORDI
	UserPointsAccount      UPA
	UserPointsLog          UPL
	Search                 SRCH
	Consultation           CN
	AppRelease             AR
	MerchantCategory       MCAT
	MembershipLevel        ML
	Node                   NODE
	PointsCfg              PCFG
	PointsRecharge         PR
	PointsSettings         PTS
	MerchantPointsSettings MPTS
	UserAddress            UA
	AgentLevel             AL
	DownlinePurchaseRecord DPR
	AgentTransactionDetail atd
	Settlement             settlement
}
