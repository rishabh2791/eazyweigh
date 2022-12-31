package main

import (
	"context"
	"eazyweigh/application"
	"eazyweigh/infrastructure/config"
	"eazyweigh/infrastructure/persistance"
	"eazyweigh/infrastructure/server"
	"eazyweigh/infrastructure/utilities"
	"eazyweigh/interfaces"
	"eazyweigh/interfaces/middlewares"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

func main() {
	logger := utilities.NewConsoleLogger()
	serverConfig := config.NewServerConfig()
	dotenvLoadError := godotenv.Load("/home/administrator/Development/code/backend/variables.env")
	if serverConfig.IsDebug() {
		dotenvLoadError = godotenv.Load("./variables.env")
	}
	if dotenvLoadError != nil {
		logger.Error(dotenvLoadError.Error())
		os.Exit(1)
	}

	repoStore, repoError := persistance.NewRepoStore(serverConfig, logger)
	if repoError != nil {
		logger.Error(repoError.Error())
		os.Exit(1)
	}

	logger.Info("Migrating Models...")
	migrationError := repoStore.Migrate()
	if migrationError != nil {
		logger.Error(migrationError.Error())
		os.Exit(1)
	}
	logger.Info("Model Migrations Done...")

	appStore := application.NewAppStore(repoStore)
	interfaceStore := interfaces.NewInterfaceStore(logger, appStore)
	middlewareStore := middlewares.NewMiddlewareStore(logger, appStore)
	httpServer := server.NewHTTPServer(*serverConfig, appStore, interfaceStore, middlewareStore)
	httpServer.Serve()

	server := http.Server{
		Addr:         ":" + serverConfig.ServerPort,
		Handler:      httpServer.Router,
		ErrorLog:     logger.StandardLogger(&hclog.StandardLoggerOptions{}),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
		IdleTimeout:  300 * time.Second,
	}

	go func() {
		logger.Info(fmt.Sprintf("Starting Server at: %s...", serverConfig.ServerAddress+":"+serverConfig.ServerPort))

		err := server.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
	sig := <-c
	logger.Info(fmt.Sprintf("Server Shutting Down... %s ", sig))

	//gracefully shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if shutDownError := server.Shutdown(ctx); shutDownError != nil {
		logger.Error(shutDownError.Error())
	}
	logger.Info("Server Shut Down.")
}
