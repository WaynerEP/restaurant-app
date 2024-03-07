package config

type SunatApisPeru struct {
	SecretToken string `mapstructure:"secret-token" json:"secret-token" yaml:"secret-token"`
	BaseURL     string `mapstructure:"base-url" json:"base-url" yaml:"base-url"`
}
