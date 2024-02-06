package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configurations struct {
	Server   ServerConfigs
	Database DatabaseConfigs
}

type ServerConfigs struct {
	Port string
}

type DatabaseConfigs struct {
	DbName     string
	DbUser     string
	DbPassword string
}

var instantiated *Configurations = nil

func GetInstance() *Configurations {
	if instantiated == nil {
		//read the configs and put it in the struct
		viper.SetConfigName("config")
		viper.AddConfigPath("./configs")
		viper.AutomaticEnv()
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
		if err := viper.Unmarshal(&instantiated); err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}

	}
	return instantiated
}
