package config

// Qiniu defines configuration parameters for Qiniu Cloud storage
type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`                                  // Storage region
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`                            // Space name
	ImgPath       string `mapstructure:"img-path" json:"img-path" yaml:"img-path"`                      // CDN acceleration domain
	AccessKey     string `mapstructure:"access-key" json:"access-key" yaml:"access-key"`                // Access Key (AK)
	SecretKey     string `mapstructure:"secret-key" json:"secret-key" yaml:"secret-key"`                // Secret Key (SK)
	UseHTTPS      bool   `mapstructure:"use-https" json:"use-https" yaml:"use-https"`                   // Whether to use HTTPS
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"use-cdn-domains" yaml:"use-cdn-domains"` // Whether to use CDN for upload acceleration
}
