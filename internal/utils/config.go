package utils

import "github.com/spf13/viper"

type Config struct {
	ApiKey       string `mapstructure:"APIKEY"`
	UploadPreset string `mapstructure:"UPLOADPRESET"`
	BaseUrl      string `mapstructure:"BASEURL"`
	CloudName    string `mapstructure:"CLOUDNAME"`
	Port         string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
