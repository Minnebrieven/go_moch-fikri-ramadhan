package domains

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBLocation string `mapstructure:"DB_LOCATION"`

	JWTKey  string `mapstructure:"JWT_KEY"`
	AppPort string `mapstructure:"APP_PORT"`
}
