package config

// Pgsql is a structure representing PostgreSQL configuration.
type Pgsql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn generates the Data Source Name (DSN) for connecting to PostgreSQL.
func (p *Pgsql) Dsn() string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + p.Dbname + " port=" + p.Port + " " + p.Config
}

// LinkDsn generates a DSN based on the provided dbname for PostgreSQL.
func (p *Pgsql) LinkDsn(dbname string) string {
	return "host=" + p.Path + " user=" + p.Username + " password=" + p.Password + " dbname=" + dbname + " port=" + p.Port + " " + p.Config
}

// GetLogMode retrieves the log mode configuration for PostgreSQL.
func (m *Pgsql) GetLogMode() string {
	return m.LogMode
}
