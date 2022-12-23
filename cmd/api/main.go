package main

import (
	"PTOBuilder/config"
	"PTOBuilder/internal/server"
	"PTOBuilder/pkg/storage"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

func main() {
	ctx := context.Background()
	err := config.GetConfigs()
	if err != nil {
		fmt.Println(err)
	}
	repo, err := storage.NewStorage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	err = repo.Ping(ctx)
	if err != nil {
		fmt.Println(err)
	}
	router := httprouter.New()
	s := server.NewServer(router)
	err = s.Run()
	if err != nil {
		fmt.Println(err)
	}
}
