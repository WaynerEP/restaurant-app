package config

// Mysql is a structure representing MySQL configuration.
type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn generates the Data Source Name (DSN) for connecting to MySQL.
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

// GetLogMode retrieves the log mode configuration for MySQL.
func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
