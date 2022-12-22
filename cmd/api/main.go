package main

import (
	"PTOBuilder/config"
	"PTOBuilder/pkg/storage"
	"context"
	"fmt"
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
}
