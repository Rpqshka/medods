package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"medods"
	"medods/pkg/handler"
	"medods/pkg/repository"
	"medods/pkg/service"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db, err := repository.NewMongoDB()
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(medods.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Printf("Medods App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("Medods App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Client().Disconnect(context.Background()); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
