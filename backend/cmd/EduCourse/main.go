package main

import (
	"Kursash/internal/config"
	"Kursash/internal/http-server/handlers"
	"Kursash/internal/http-server/middleware"
	"Kursash/internal/repository"
	"Kursash/internal/service"
	"Kursash/internal/storage/PostgreSQL"
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/go-ozzo/ozzo-log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// @title           EduCourse
// @version         1.0
// @host      localhost:8080

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	logger := setupLogger(cfg.LogsPath)
	logger.Open()
	defer logger.Close()
	storage, err := PostgreSQL.NewStorage(context.TODO(), cfg.StoragePath)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("Storage connected successfully")
	repo := repository.NewDB(storage)
	router := gin.Default()
	CORSconfig := cors.DefaultConfig()
	CORSconfig.AllowOrigins = []string{"*"}
	CORSconfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	router.Use(cors.New(CORSconfig))
	authMW, err := middleware.GetAuthMW(cfg.Username, cfg.Password)
	if err != nil {
		logger.Error("Error while getting auth middleware", "error", err)
		os.Exit(1)
	}
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	serv := service.NewService(repo)
	handler := handlers.NewHandler(router, serv, logger, authMW)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	server := &http.Server{
		Addr:         cfg.Address,
		Handler:      handler,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Error("failed to start server: %s", err.Error())
		}
	}()

	logger.Info("Server started successfully")
	<-signalChan

	logger.Info("Shutting down the server")
}

func setupLogger(logsPath string) *log.Logger {
	logger := log.NewLogger()

	t1 := log.NewConsoleTarget()
	t2 := log.NewFileTarget()
	t2.FileName = logsPath
	logger.Targets = append(logger.Targets, t1, t2)

	return logger
}
