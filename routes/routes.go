package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoute(r)
	CountryRoute(r)
	LogRegRoutes(r)
	TripRoute(r)
	TransaksiRoute(r)
}
