package configs

var cfg *Config

type (
	Config struct {
		DB DBConfig
	}

	DBConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
)

func init() {
	cfg = new(Config)
	cfg.DB.Host = "localhost"
	cfg.DB.Port = "5432"
	cfg.DB.User = "user"
	cfg.DB.Password = "password123"
	cfg.DB.Database = "groceryDB"
}

func GetDB() DBConfig {
	return cfg.DB
}
