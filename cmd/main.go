package main

import (
	"log"

	"github.com/EvgeniyGnetnev/GenreID"
	"github.com/EvgeniyGnetnev/GenreID/pkg/handler"
	"github.com/EvgeniyGnetnev/GenreID/pkg/repository"
	"github.com/EvgeniyGnetnev/GenreID/pkg/service"
	"github.com/spf13/viper"
)

func main(){
	if err := initConfig(); err != nil {
		log.Fatalf("error intializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(genreid.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}