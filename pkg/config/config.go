package config

import "github.com/spf13/viper"

type Config struct {
	App struct {
		Name string
		Port int
	}
}

func Load(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	err := viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
