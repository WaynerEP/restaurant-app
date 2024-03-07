package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // Recipient: Multiple recipients separated by commas
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // Port
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // Sender
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // Server address
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // Whether SSL
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // Secret key
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // Nickname
}
