package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`                                  // Environment value
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`                               // Port value
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // Database type: mysql(default)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`                   // Oss type
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // Multiple login interception
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                // Use redis
	LimitCountIP  int    `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int    `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
}
