package routes

import (
	"net/http"

	"github.com/andersonluizneto/storeGH/internal/customer/interfaces/http/handler"
)

func RegisterRoutes(mux *http.ServeMux, handler *handler.CustomerHandler) {
	mux.HandleFunc("POST /customers", handler.Create)
	mux.HandleFunc("GET /customers", handler.FindAll)
	mux.HandleFunc("GET /customers", handler.FindByID)
	mux.HandleFunc("PUT /customers", handler.Update)
	mux.HandleFunc("DELETE /customers", handler.Delete)
}
