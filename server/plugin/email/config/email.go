package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // Recipients: Multiple recipients separated by commas, e.g., a@qq.com b@qq.com. In formal development, please use this as a parameter.
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // Sender: Your own email address from which you want to send emails.
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // Server address: For example, smtp.qq.com. Please check the SMTP protocol for QQ or the email service you want to use.
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // Secret: The key used for login. It's recommended not to use the email password. Apply for a key for login from the SMTP settings of your email provider.
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // Nickname: Sender's nickname, usually your own email address.
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // Port: Please check the SMTP protocol for QQ or the email service you want to use. Often 465.
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // IsSSL: Whether to enable SSL.
}
