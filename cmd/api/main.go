package main

import (
	"PTOBuilder/config"
	"PTOBuilder/internal/server"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()
	ctx := context.Background()
	err := config.GetConfigs()
	if err != nil {
		logger.Fatal(err)
	}
	repo, err := storage.NewStorage(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	err = repo.Ping(ctx)
	if err != nil {
		logger.Fatal(err)
	}
	router := httprouter.New()
	s := server.NewServer(router)
	err = s.Run()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("all done right")
}
