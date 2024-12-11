package initializers

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AutomaticEnv()

	viper.BindEnv("POSTGRES_HOST")
	viper.BindEnv("POSTGRES_USER")
	viper.BindEnv("POSTGRES_PASSWORD")
	viper.BindEnv("POSTGRES_DB")
	viper.BindEnv("POSTGRES_PORT")
	viper.BindEnv("PORT")
	viper.BindEnv("CLIENT_ORIGIN")
	viper.BindEnv("ACCESS_TOKEN_PRIVATE_KEY")
	viper.BindEnv("ACCESS_TOKEN_PUBLIC_KEY")
	viper.BindEnv("REFRESH_TOKEN_PRIVATE_KEY")
	viper.BindEnv("REFRESH_TOKEN_PUBLIC_KEY")
	viper.BindEnv("ACCESS_TOKEN_EXPIRED_IN")
	viper.BindEnv("REFRESH_TOKEN_EXPIRED_IN")
	viper.BindEnv("ACCESS_TOKEN_MAXAGE")
	viper.BindEnv("REFRESH_TOKEN_MAXAGE")

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
