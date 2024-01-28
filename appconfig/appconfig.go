package appconfig

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type AuthConfig struct {
	SHASalt       string
	JWTSigningKey string
	JWTTokenTTL   int
}

type FullConfig struct {
	Host       string
	DBConfig   DBConfig
	AuthConfig AuthConfig
}

var config *FullConfig = nil

func Get() (config *FullConfig, err error) {
	if config == nil {
		if err := initConfigFile(); err != nil {
			return config, err
		}
		config = new(FullConfig)
		config.Host = viper.GetString("port")
		config.DBConfig = DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetString("db.port"),
			Username: viper.GetString("db.username"),
			DBName:   viper.GetString("db.dbname"),
			SSLMode:  viper.GetString("db.sslmode"),
			Password: viper.GetString("db.password"),
		}
		config.AuthConfig = AuthConfig{
			SHASalt:       viper.GetString("auth.sha_salt"),
			JWTSigningKey: viper.GetString("auth.jwt_signing_key"),
			JWTTokenTTL:   viper.GetInt("auth.jwt_token_ttl"),
		}
	}

	return config, nil
}

func initConfigFile() error {
	viper.AddConfigPath("appconfig")
	viper.SetConfigName("appconfig")
	return viper.ReadInConfig()
}
