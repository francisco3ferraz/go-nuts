package router

import (
	"net/http"

	"github.com/francisco3ferraz/go-nuts/internal/api/handler"
	"github.com/francisco3ferraz/go-nuts/internal/auth"
	"github.com/francisco3ferraz/go-nuts/internal/orders"
)

func New() http.Handler {
	mux := http.NewServeMux()
	authMiddleware := auth.NewMiddleware()

	ordersRepository := orders.NewInMemoryRepository()
	ordersService := orders.NewService(ordersRepository)
	ordersHandler := handler.NewOrdersHandler(ordersService)

	mux.HandleFunc("GET /healthz", handler.Health)
	mux.Handle("GET /v1/ping", authMiddleware.RequireAPIKey(http.HandlerFunc(handler.Ping)))
	mux.Handle("GET /v1/orders", authMiddleware.RequireAPIKey(http.HandlerFunc(ordersHandler.List)))

	return mux
}
