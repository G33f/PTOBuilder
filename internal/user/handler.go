package user

import (
	"PTOBuilder/internal/handlers"
	"PTOBuilder/internal/user/model"
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
	router.Post("/user/create", h.UserRegistration)
}

func (h *handler) UserRegistration(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	ctx := context.Background()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	err = h.useCase.UserRegistration(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.log.Info("User has created!")
}
