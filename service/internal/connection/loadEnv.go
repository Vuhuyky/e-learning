package connection

import (
	"log"

	"github.com/spf13/viper"
)

func loadYml() {
	viper.SetConfigName("service")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error read file config: %v", err)
	}

	if err := viper.Unmarshal(&conn); err != nil {
		log.Fatalf("error map to struct: %v", err)
	}
}
