package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBHost string `mapstructure:"db_host"`
	DBName string `mapstructure:"db_name"`
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %s \n", err)
		return
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %s\n", err)
		return
	}

	fmt.Printf("DB User: %s \n", config.DBUser)
	fmt.Printf("DB Password: %s \n", config.DBPass)
	fmt.Printf("DB Host: %s \n", config.DBHost)
	fmt.Printf("DB Name: %s \n", config.DBName)
}
