package main

import (
	"API"
	"API/pkg/handler"
	"API/pkg/repository"
	"API/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("failed to initialise configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environmental variables: %s", err.Error())
	}
	db, err := repository.NewSQLDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		//SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialise database: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(API.Server)
	if err := srv.Run(viper.GetString(viper.GetString("8000")), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
