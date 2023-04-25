package character

import (
	"PTOBuilder/internal/character/model"
	"PTOBuilder/internal/handlers"
	"PTOBuilder/pkg/logging"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type handler struct {
	log     *logging.Logger
	useCase UseCase
}

func NewHandler(log *logging.Logger, useCase UseCase) handlers.Handler {
	return &handler{
		log:     log,
		useCase: useCase,
	}
}

func (h *handler) MainRoutsHandler(router chi.Router) {
	router.Post("/Character/Hero/Create", h.CreateCharacter)
	router.Post("/Character/Role/Create", h.CreateRole)
	router.Get("/Character/Get", h.GetCharacter)
}

func (h *handler) CreateRole(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	role := model.Role{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&role)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	if err = h.useCase.CreateRole(ctx, &role); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.log.Info("role CreateRole work right")
}

func (h *handler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hero := model.Character{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&hero)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	if err = h.useCase.CreateCharacter(ctx, &hero); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.log.Info("hero CreateRole work right")
}

func (h *handler) GetCharacter(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	hero, err := h.useCase.GetCharacter(ctx, "Feng")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(hero)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.log.Info("hero GetCharacter work right")
}
