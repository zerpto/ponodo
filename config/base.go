package config

type DB struct {
	Host     string `mapstructure:"DB_HOST"`
	Name     string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Port     string `mapstructure:"DB_PORT"`
}

type Config struct {
	App   string `mapstructure:"APP"`
	Env   string `mapstructure:"ENV"`
	Debug bool   `mapstructure:"DEBUG"`

	HTTPPort  string `mapstructure:"HTTP_PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`

	DB DB
}
