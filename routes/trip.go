package routes

import (
	"Backend/handlers"
	"Backend/pkg/middleware"
	"Backend/pkg/mysql"
	"Backend/repositories"

	"github.com/gorilla/mux"
)

func TripRoute(r *mux.Router) {
	tripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandleTrip(tripRepository)

	r.HandleFunc("/trips", h.FindTrip).Methods("GET")
	r.HandleFunc("/trip/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.UploadFile(h.CreateTrip)).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.UploadFile(h.UpdatedTrip)).Methods("PATCH")
	r.HandleFunc("/trip/{id}", h.DeleteTrip).Methods("Delete")
}
