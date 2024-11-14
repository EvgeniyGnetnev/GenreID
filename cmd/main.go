package main

import (
	"github.com/EvgeniyGnetnev/GenreID"
	"github.com/spf13/viper"
	"log"
)

func main(){
	if err := initConfig(); err != nil {
		log.Fatalf("error intializing configs: %s", err.Error())
	}

	srv := new(genreid.Server)
	if err := srv.Run(viper.GetString("port")); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}