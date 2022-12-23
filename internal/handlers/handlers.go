package handlers

import "github.com/julienschmidt/httprouter"

type Handler interface {
	MainRoutsHandler(router *httprouter.Router)
}
