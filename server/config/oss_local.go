package config

type Local struct {
	Path      string `mapstructure:"path" json:"path" yaml:"path"`                   // local file access path
	StorePath string `mapstructure:"store-path" json:"store-path" yaml:"store-path"` // local file repository path
}
