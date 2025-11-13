package service

var Service = new(service)

type service struct {
	AppUsers             appUsers
	Merchants            mc
	Banner               banner
	MerchantAnnouncement MA
	MerchantAdmin        MADM
	ProductCategory      PC
	Product              P
	ProductSku           SKU
	ProductSpec          PS
	ProductSpecOption    PSO
}
