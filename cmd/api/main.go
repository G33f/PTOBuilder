package main

import (
	"PTOBuilder/config"
	"PTOBuilder/internal/server"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"context"
	"github.com/julienschmidt/httprouter"
)

//TODO DEL THIS
//
//func tmp() {
//	a, _ := calculator.Calculate(fmt.Sprintf("%f + %f", 1.2, 1.3))
//	fmt.Println(a)
//}
//----------

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
	//characterHandler := character.NewHandler(&logger)
	//characterHandler.MainRoutsHandler(router)
	s := server.NewServer(router)
	err = s.Run()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("all done right")
}
