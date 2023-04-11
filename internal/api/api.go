package api

import (
	"PTOBuilder/internal/character"
	characterR "PTOBuilder/internal/character/repo"
	characterUC "PTOBuilder/internal/character/usecase"
	"PTOBuilder/internal/server"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	apiServer *server.Server
	log       *logging.Logger
	repo      *pgxpool.Pool
}

func NewAPI(log *logging.Logger) *API {
	return &API{
		log: log,
	}
}

func (api *API) Init() {
	var err error

	api.log.Info("API initialization...")

	// DataBase connection and creating client interface
	ctx := context.Background()
	api.log.Info("Connection to storage...")
	api.repo, err = storage.NewStorage(ctx, api.log)
	if err != nil {
		api.log.Fatal(err)
	}

	// Server creation
	api.log.Info("Creating Server...")
	router := httprouter.New()
	api.apiServer = server.NewServer(router)

	// Creating handler, interfaces UseCase and repository for character
	api.log.Info("Initialising character...")
	characterRepo := characterR.NewRepo(api.log, api.repo)
	characterUseCase := characterUC.NewUseCase(api.log, &characterRepo)
	characterHandler := character.NewHandler(api.log, &characterUseCase)
	characterHandler.MainRoutsHandler(router)
}

func (api *API) Start() {
	if err := api.apiServer.Run(); err != nil {
		api.log.Fatal(err, "An error occurred while running the api")
	}
}