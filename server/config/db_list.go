package config

// DsnProvider is an interface for providing a DSN (Data Source Name).
type DsnProvider interface {
	Dsn() string
}

// GeneralDB is a structure representing general database configuration.
type GeneralDB struct {
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`       // Advanced configuration
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`    // Database name
	Username     string `mapstructure:"username" json:"username" yaml:"username"` // Database username
	Password     string `mapstructure:"password" json:"password" yaml:"password"` // Database password
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        // Database engine, default is InnoDB
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // Whether to enable global Gorm logging
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // Maximum number of idle connections
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // Maximum number of open connections to the database
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   // Whether to globally disable plural forms, true means enabled
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // Whether to write logs to files via zap
}

// SpecializedDB is a structure representing specialized database configuration.
type SpecializedDB struct {
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	AliasName string `mapstructure:"alias-name" json:"alias-name" yaml:"alias-name"`
	GeneralDB `yaml:",inline" mapstructure:",squash"`
	Disable   bool `mapstructure:"disable" json:"disable" yaml:"disable"`
}
