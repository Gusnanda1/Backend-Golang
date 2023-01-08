package routes

import (
	"Backend/handlers"
	"Backend/pkg/middleware"
	"Backend/pkg/mysql"
	"Backend/repositories"

	"github.com/gorilla/mux"
)

func TransaksiRoute(r *mux.Router) {
	TransaksiRepository := repositories.RepositoryTransaksi(mysql.DB)
	h := handlers.HandlerTransaksi(TransaksiRepository)

	r.HandleFunc("/transaksi", h.FindTransaksi).Methods("GET")
	r.HandleFunc("/transaksi/{id}", middleware.Auth(h.GetTransaksi)).Methods("GET")
	r.HandleFunc("/transaksi", middleware.Auth(h.AddTransaksi)).Methods("POST")
	// r.HandleFunc("/transaksi/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTransaksi))).Methods("PATCH")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
	r.HandleFunc("/transaksi/{id}", middleware.Auth(h.DeleteTransaksi)).Methods("Delete")
}
