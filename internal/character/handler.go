package character

import (
	"PTOBuilder/internal/character/model"
	"PTOBuilder/internal/handlers"
	"PTOBuilder/pkg/logging"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type handler struct {
	log *logging.Logger
}

func NewHandler(l *logging.Logger) handlers.Handler {
	return &handler{l}
}

func (h *handler) MainRoutsHandler(router *httprouter.Router) {
	router.POST("/Character/Hero/Create", h.CreateCharacter)
	router.POST("/Character/Role/Create", h.CreateRole)
	router.GET("/Character/Get", h.GetCharacters)
}

func (h *handler) CreateRole(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	role := model.Role{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&role)
	if err != nil {
		w.Write([]byte("cannot unmarshal body"))
		w.WriteHeader(400)
		h.log.Info(err)
	}
	w.WriteHeader(200)
	h.log.Info("rout CreateRole work right")
}

func (h *handler) CreateCharacter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create character"))
	w.WriteHeader(200)
	h.log.Info("rout CreateCharacter work right")
}

func (h *handler) GetCharacters(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is get character"))
	w.WriteHeader(200)
	h.log.Info("rout GetCharacter work right")
}
