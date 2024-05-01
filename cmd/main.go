package main

import (
	"github.com/halimonalexander/todo"
	"github.com/halimonalexander/todo/pkg/handler"
	"github.com/halimonalexander/todo/pkg/repository"
	"github.com/halimonalexander/todo/pkg/service"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error occurred while loading config file: #{err.Error()}")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE_NAME"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("error occurred while creating DB connection: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	servs := service.NewService(repos)
	handls := handler.NewHandler(servs)
	//handls := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handls.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
