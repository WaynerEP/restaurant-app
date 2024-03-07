package config

// Captcha defines configuration parameters for captcha settings
type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // Length of the captcha key
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // Width of the captcha image
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // Height of the captcha image
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // Enable anti-brute-force captcha; 0 means captcha required for every login, other numbers represent the number of incorrect password attempts before captcha appears
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // Timeout for anti-brute-force captcha, in seconds
}
