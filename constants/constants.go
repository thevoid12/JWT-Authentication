package constant

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func GetEnvConstFromViper(key string) string {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		fmt.Println(err)
		log.Fatalf("unable to read .env file from viper", err)
		return ""
	}
	return fmt.Sprintf("%v", viper.Get(key))
}

func SetEnvConstIntoViper(key string, value string) {
	viper.SetConfigFile(".env")
	viper.AddConfigPath("../")
	viper.Set(key, value)
}
