package auth

import (
	"PTOBuilder/internal/auth/model"
	"PTOBuilder/internal/handlers"
	"PTOBuilder/pkg/logging"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync"
	"time"
)

const (
	admin = "admin"
	user  = "user"
)

type ValidationFunctions struct {
	CheckAdmin func(string) bool
	CheckAuth  func(string) bool
}

type handler struct {
	log     *logging.Logger
	useCase UseCase
	mu      *sync.RWMutex
	jwt     map[string]model.User
}

func NewHandler(log *logging.Logger, useCase UseCase, mu *sync.RWMutex) (handlers.Handler, ValidationFunctions) {
	h := handler{
		log:     log,
		useCase: useCase,
		mu:      mu,
		jwt:     map[string]model.User{},
	}

	validationFunctions := ValidationFunctions{
		CheckAdmin: h.CheckAdmin,
		CheckAuth:  h.CheckAuth,
	}
	go h.tokensTracking()
	return &h, validationFunctions
}

func (h *handler) MainRoutsHandler(router chi.Router) {
	router.Post("/auth/signUp", h.SignUp)
	router.Post("/auth/signIn", h.SignIn)
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	ctx := context.Background()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	err = h.useCase.SignUp(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	h.log.Info("User has created!")
}

func (h *handler) SignIn(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	ctx := context.Background()
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	token, err := h.useCase.SignIn(ctx, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		h.log.Info(err)
		return
	}
	h.mu.Lock()
	h.jwt[token] = user
	h.mu.Unlock()
	w.Header().Add("Authorization", fmt.Sprintf("Bearer %s", token))
	w.WriteHeader(http.StatusOK)
	fmt.Println(token)
	h.log.Info("User logIn!")
	fmt.Println(h.jwt)
}

func (h *handler) CheckAuth(token string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if _, ok := h.jwt[token]; ok {
		return true
	}
	return false
}

func (h *handler) CheckAdmin(token string) bool {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.jwt[token].Role == admin {
		return true
	}
	return false
}

func (h *handler) tokensTracking() {
	for {
		time.Sleep(15 * time.Minute)
		h.mu.Lock()
		for k, _ := range h.jwt {
			if err := h.useCase.ValidToken(k); err != nil {
				delete(h.jwt, k)
			}
		}
		h.mu.Unlock()
	}
}
