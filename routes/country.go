package routes

import (
	"Backend/handlers"
	"Backend/pkg/middleware"
	"Backend/pkg/mysql"
	"Backend/repositories"

	"github.com/gorilla/mux"
)

func CountryRoute(r *mux.Router) {
	countryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandleCountry(countryRepository)

	r.HandleFunc("/country", h.FindCountry).Methods("GET")
	r.HandleFunc("/country/{id}", h.GetCountry).Methods("GET")
	r.HandleFunc("/country", middleware.Auth(h.AddCountry)).Methods("POST")
	r.HandleFunc("/country/{id}", middleware.Auth(h.UpdateCountry)).Methods("PATCH")
	r.HandleFunc("/country/{id}", middleware.Auth(h.DeleteCountry)).Methods("Delete")
}
