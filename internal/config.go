package internal

import "github.com/spf13/viper"

type Config struct {
	DatabaseURL string
}

func GetConfig() *Config {
	viper.AutomaticEnv()
	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.ReadInConfig()

	//Note: Viper configuration keys are case insensitive.
	viper.SetDefault("database_url", "some_random_default_value")

	config := Config{
		DatabaseURL: viper.GetString("DATABASE_URL"),
	}

	return &config
}
