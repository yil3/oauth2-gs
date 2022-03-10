package health

import (
	"github.com/gorilla/mux"
	"oauth2-gs/util/routes"
)

// ServiceInterface defines exported methods
type ServiceInterface interface {
	// Exported methods
	GetRoutes() []routes.Route
	RegisterRoutes(router *mux.Router, prefix string)
	Close()
}
