package config

type AppConfig struct {
	Name string
	Env  string
	Host string
	Port string
}

func LoadAppConfig() AppConfig {
	return AppConfig{
		Name: getEnv("APP_NAME", "boilerplate-one"),
		Env:  getEnv("APP_ENV", "development"),
		Host: getEnv("APP_HOST", "localhost"),
		Port: getEnv("APP_PORT", "3000"),
	}
}
