package config

type DBConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	Name               string
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime int
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:               getEnv("DB_HOST", "localhost"),
		Port:               getEnv("DB_PORT", "5432"),
		User:               getEnv("DB_USER", "postgres"),
		Password:           getEnv("DB_PASSWORD", "postgres"),
		Name:               getEnv("DB_NAME", "postgres"),
		SetMaxIdleConns:    getEnvInt("DB_SET_MAX_IDLE_CONNS", 10),
		SetMaxOpenConns:    getEnvInt("DB_SET_MAX_OPEN_CONNS", 100),
		SetConnMaxLifetime: getEnvInt("DB_SET_CONN_MAX_LIFETIME", 30),
	}
}
