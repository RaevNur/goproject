package api

import (
	"net/http"

	"forum/internal/service"
)

type ApiHandler struct {
	service *service.Service
}

func NewApiHandler(service *service.Service) *ApiHandler {
	return &ApiHandler{
		service: service,
	}
}

func (v *ApiHandler) InitRoutes(mux *http.ServeMux) {
	// Init Api routes
}
