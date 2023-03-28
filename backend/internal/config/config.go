package config

import "github.com/spf13/viper"

type Config struct {
	DbName               string `mapstructure:"db_name"`
	DefaultTaskListTitle string `mapstructure:"default_task_list_title"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
