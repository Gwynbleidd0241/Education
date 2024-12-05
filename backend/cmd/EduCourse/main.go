package main

import (
	backend "Kursash"
	"Kursash/internal/http-server/handlers"
	"Kursash/internal/repository"
	"Kursash/internal/service"
	"Kursash/notifications"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// @title           EduCourse
// @version         1.0
// @host      localhost:8080

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	router := gin.Default()
	CORSconfig := cors.DefaultConfig()
	CORSconfig.AllowOrigins = []string{"http://google.com", "http://localhost:2001"}
	CORSconfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	CORSconfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	router.Use(cors.New(CORSconfig))

	mailSender := notifications.NewMailSender(
		viper.GetString("email.host"),
		viper.GetInt("email.port"),
		viper.GetString("email.username"),
		viper.GetString("email.password"),
	)

	repos := repository.NewRepository(db)
	services := service.NewService(repos, service.NewNotifications(mailSender))
	handlers := handlers.NewHandler(services)

	srv := new(backend.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes(router)); err != nil {
		logrus.Fatal("error occured while running hhtp server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
