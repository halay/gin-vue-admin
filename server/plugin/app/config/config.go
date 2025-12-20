package config

type Config struct {
	AuthorityId  uint     `mapstructure:"authority-id" json:"authority-id" yaml:"authority-id"`
	YApiUrl      string   `mapstructure:"y-api-url" json:"y-api-url" yaml:"y-api-url"`
	YApiTimeout  int      `mapstructure:"y-api-time-out" json:"y-api-time-out" yaml:"y-api-time-out"`
	XCgProApiKey string   `mapstructure:"x-cg-pro-api-key" json:"x-cg-pro-api-key" yaml:"x-cg-pro-api-key"`
	XCgProApiUrl string   `mapstructure:"x-cg-pro-api-url" json:"x-cg-pro-api-url" yaml:"x-cg-pro-api-url"`
	HrefUrl      string   `mapstructure:"href-url" json:"href-url" yaml:"href-url"`
	XCgApiIds    []string `mapstructure:"x-cg-api-ids" json:"x-cg-api-ids" yaml:"x-cg-api-ids"` // bitcoin,ethereum,ripple,binancecoin,solana,tron,staked-ether,dogecoin
}
