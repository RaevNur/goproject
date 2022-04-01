package handler

import (
	"fmt"
	"net/http"

	"forum/internal/service"
	"forum/internal/web/handler/api"
	"forum/internal/web/handler/view"
)

type MainHandler struct {
	apiHandler  *api.ApiHandler
	viewHandler *view.ViewHandler
}

func NewMainHandler(service *service.Service) (*MainHandler, error) {
	apiHandler := api.NewApiHandler(service)
	viewHandler, err := view.NewViewHandler()
	if err != nil {
		return nil, fmt.Errorf("NewMainHandler: %w", err)
	}
	return &MainHandler{
		apiHandler:  apiHandler,
		viewHandler: viewHandler,
	}, nil
}

func (m *MainHandler) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	m.apiHandler.InitRoutes(mux)
	m.viewHandler.InitRoutes(mux)
	return mux
}
