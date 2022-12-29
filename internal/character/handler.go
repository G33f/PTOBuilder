package character

import (
	"PTOBuilder/internal/handlers"
	"PTOBuilder/pkg/logging"
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
	router.POST("/Character/Create", h.CreateCharacter)
	router.GET("/Character/Get", h.GetCharacter)
}

func (h *handler) CreateCharacter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is create character"))
	w.WriteHeader(200)
	h.log.Info("rout CreateCharacter work right")
}

func (h *handler) GetCharacter(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is get character"))
	w.WriteHeader(200)
	h.log.Info("rout GetCharacter work right")
}
