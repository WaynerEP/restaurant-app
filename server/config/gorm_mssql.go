package config

// Mssql is a structure representing Microsoft SQL Server (MSSQL) configuration.
type Mssql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn generates the Data Source Name (DSN) for connecting to MSSQL.
// Dsn dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm&encrypt=disable"
func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
}

// GetLogMode retrieves the log mode configuration for MSSQL.
func (m *Mssql) GetLogMode() string {
	return m.LogMode
}
