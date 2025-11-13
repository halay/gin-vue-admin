package config

type Config struct {
	AuthorityId uint   `mapstructure:"authority-id" json:"authority-id" yaml:"authority-id"`
	YApiUrl     string `mapstructure:"y-api-url" json:"y-api-url" yaml:"y-api-url"`
	YApiTimeout int    `mapstructure:"y-api-time-out" json:"y-api-time-out" yaml:"y-api-time-out"`
}
