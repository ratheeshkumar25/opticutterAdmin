package config

import "github.com/spf13/viper"

type Config struct {
	SECERETKEY    string `mapstructure:"JWTKEY"`
	Host          string `mapstructure:"HOST"`
	User          string `mapstructure:"USER"`
	Password      string `mapstructure:"PASSWORD"`
	Database      string `mapstructure:"DBNAME"`
	Port          string `mapstructure:"PORT"`
	Sslmode       string `mapstructure:"SSL"`
	GrpcUserPort  string `mapstructure:"GRPCUSERPORT"`
	AdminPort     string `mapstructure:"GRPCADMINPORT"`
	MaterialPort  string `mapstructure:"GRPCMATERIALPORT"`
	AdminEmail    string `mapstructure:"ADMINEMAIL"`
	AdminPassword string `mapstructure:"ADMINPASSWORD"`
}

func LoadConfig() *Config {
	var config Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&config)
	return &config
}
