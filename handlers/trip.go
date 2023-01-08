package handlers

import (
	dto "Backend/dto/result"
	tripdto "Backend/dto/trip"
	"Backend/models"
	"Backend/repositories"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type handleTrip struct {
	TripRepository repositories.TripRepository
}

func HandleTrip(TripRepository repositories.TripRepository) *handleTrip {
	return &handleTrip{TripRepository}
}

func (h *handleTrip) FindTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	trip, err := h.TripRepository.FindTrip()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode((err.Error()))
	}
	for i, p := range trip {
		trip[i].Image = path_file + p.Image
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)

}

func (h *handleTrip) GetTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	trip.Image = path_file + trip.Image
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: trip}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTrip) CreateTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	kuota, _ := strconv.Atoi(r.FormValue("kuota"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	request := tripdto.TripRequest{
		Title:        r.FormValue("title"),
		Country_id:   country_id,
		Accomodation: r.FormValue("accomodation"),
		Transport:    r.FormValue("transport"),
		Eat:          r.FormValue("eat"),
		Day:          day,
		Night:        night,
		Date:         date,
		Price:        price,
		Kuota:        kuota,
		Description:  r.FormValue("description"),
		Image:        filename,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	trip := models.Trip{
		Title:        request.Title,
		Country_id:   request.Country_id,
		Accomodation: request.Accomodation,
		Transport:    request.Transport,
		Eat:          request.Eat,
		Day:          request.Day,
		Night:        request.Night,
		Date:         request.Date,
		Price:        request.Price,
		Kuota:        request.Kuota,
		Description:  request.Description,
		Image:        request.Image,
	}

	data, err := h.TripRepository.Createtrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	test, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTrip) UpdatedTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	country_id, _ := strconv.Atoi(r.FormValue("country_id"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	night, _ := strconv.Atoi(r.FormValue("night"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	kuota, _ := strconv.Atoi(r.FormValue("kuota"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	request := tripdto.TripUpdateRequest{
		Title:        r.FormValue("title"),
		Country_id:   country_id,
		Accomodation: r.FormValue("accomodation"),
		Transport:    r.FormValue("transport"),
		Eat:          r.FormValue("eat"),
		Day:          day,
		Night:        night,
		Date:         date,
		Price:        price,
		Kuota:        kuota,
		Description:  r.FormValue("description"),
		Image:        filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Title != "" {
		trip.Title = request.Title
	}

	if request.Country_id != 0 {
		trip.Country_id = request.Country_id
	}

	if request.Accomodation != "" {
		trip.Accomodation = request.Accomodation
	}

	if request.Transport != "" {
		trip.Transport = request.Transport
	}

	if request.Eat != "" {
		trip.Eat = request.Eat
	}

	if request.Day != 0 {
		trip.Day = request.Day
	}

	if request.Night != 0 {
		trip.Night = request.Night
	}

	time := time.Now()
	if request.Date != time {
		trip.Date = request.Date
	}

	if request.Price != 0 {
		trip.Price = request.Price
	}

	if request.Kuota != 0 {
		trip.Kuota = request.Kuota
	}

	if request.Description != "" {
		trip.Description = request.Description
	}

	if request.Image != "" {
		trip.Image = request.Image
	}

	data, err := h.TripRepository.UpdatedTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	test, err := h.TripRepository.GetTrip(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTrip) DeleteTrip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	trip, err := h.TripRepository.GetTrip(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TripRepository.DeleteTrip(trip)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTripDel(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseTripDel(u models.Trip) tripdto.TripResponsedel {
	return tripdto.TripResponsedel{
		ID: u.ID,
	}
}
