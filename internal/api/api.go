package api

import (
	"PTOBuilder/internal/character"
	characterR "PTOBuilder/internal/character/repo"
	characterUC "PTOBuilder/internal/character/usecase"
	"PTOBuilder/internal/server"
	"PTOBuilder/internal/user"
	userR "PTOBuilder/internal/user/repo"
	userUC "PTOBuilder/internal/user/usecase"
	"PTOBuilder/pkg/logging"
	"PTOBuilder/pkg/storage"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
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
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	api.apiServer = server.NewServer(router)

	// Creating handler, interfaces UseCase and repository for character
	api.log.Info("Initialising character...")

	characterRepo := characterR.NewRepo(api.log, api.repo)
	characterUseCase := characterUC.NewUseCase(api.log, characterRepo)
	characterHandler := character.NewHandler(api.log, characterUseCase)

	api.log.Info("character initialised")

	api.log.Info("Initialising user...")

	userRepo := userR.NewRepo(api.log, api.repo)
	userUserCase := userUC.NewUseCase(api.log, userRepo)
	userHandler := user.NewHandler(api.log, userUserCase)

	api.log.Info("user initialised")

	characterHandler.MainRoutsHandler(router)
	userHandler.MainRoutsHandler(router)
}

func (api *API) Start() {
	if err := api.apiServer.Run(); err != nil {
		api.log.Fatal(err, "An error occurred while running the api")
	}
}
