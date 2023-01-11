package handlers

import (
	dto "Backend/dto/result"
	transaksidto "Backend/dto/transaksi"
	"Backend/models"
	"Backend/repositories"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
	"gopkg.in/gomail.v2"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var path_file = "http://localhost:7000/uploads/"

type handleTransaksi struct {
	TransaksiRepository repositories.TransaksiRepository
}

func HandlerTransaksi(TransaksiRepository repositories.TransaksiRepository) *handleTransaksi {
	return &handleTransaksi{TransaksiRepository}
}

func (h *handleTransaksi) FindTransaksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transaksi, err := h.TransaksiRepository.FindTransaksi()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	for i, p := range transaksi {
		transaksi[i].Image = path_file + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaksi}
	json.NewEncoder(w).Encode(response)

}

func (h *handleTransaksi) GetTransaksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	transaksi, err := h.TransaksiRepository.GetTransaksi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: transaksi}
	json.NewEncoder(w).Encode(response)
}

func (h *handleTransaksi) AddTransaksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get data user token
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	// dataContex := r.Context().Value("dataFile") // add this code
	// filename := dataContex.(string)             // add this code

	counterqty, _ := strconv.Atoi(r.FormValue("counter_qty"))
	total, _ := strconv.Atoi(r.FormValue("total"))
	tripid, _ := strconv.Atoi(r.FormValue("trip_id"))
	// userid, _ := strconv.Atoi(r.FormValue("user_id"))

	request := transaksidto.CreateTransaksiRequest{
		CounterQTY: counterqty,
		Total:      total,
		Status:     "pending",
		// Image:      filename,
		Trip_id: tripid,
		UserID:  userId,
	}
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var TransIdIsMatch = false
	var TransactionId int
	for !TransIdIsMatch {
		TransactionId = rand.Intn(10000) - rand.Intn(100)
		transactionData, _ := h.TransaksiRepository.GetTransaksi(TransactionId)
		if transactionData.ID == 0 {
			TransIdIsMatch = true
		}
	}

	transaksi := models.Transaction{
		ID:         TransactionId,
		CounterQTY: request.CounterQTY,
		Total:      request.Total,
		Status:     request.Status,
		Image:      request.Image,
		Trip_id:    request.Trip_id,
		UserID:     request.UserID,
	}

	data, err := h.TransaksiRepository.AddTransaksi(transaksi)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	bagus, err := h.TransaksiRepository.GetTransaksi(data.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	var s = snap.Client{}
	s.New("SB-Mid-server-3PxUPSmyBSouHrnBvXxzDHAv", midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(bagus.ID),
			GrossAmt: int64(bagus.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: bagus.User.Fullname,
			Email: bagus.User.Email,
		},
	}

	// 3. Execute request create Snap transaction to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: snapResp}
	json.NewEncoder(w).Encode(response)
}

// func (h *handleTransaksi) UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	dataContex := r.Context().Value("dataFile") // add this code
// 	filename := dataContex.(string)             // add this code

// 	counterqty, _ := strconv.Atoi(r.FormValue("counter_qty"))
// 	total, _ := strconv.Atoi(r.FormValue("total"))
// 	request := transaksidto.UpdateTransaksiRequest{
// 		CounterQTY: counterqty,
// 		Total:      total,
// 		Status:     r.FormValue("status"),
// 		Image:      filename,
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])
// 	transaksi, err := h.TransaksiRepository.GetTransaksi(int(id))
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(w).Encode(err.Error())
// 		return
// 	}

// 	//counter qty
// 	if request.CounterQTY != 0 {
// 		transaksi.CounterQTY = request.CounterQTY
// 	}

// 	//total
// 	if request.Total != 0 {
// 		transaksi.Total = request.Total
// 	}

// 	//status
// 	if request.Status != "" {
// 		transaksi.Status = request.Status
// 	}

// 	//image
// 	if request.Image != "" {
// 		transaksi.Image = request.Image
// 	}

// 	//trip id
// 	tripid, _ := strconv.Atoi(r.FormValue("trip_id"))
// 	if tripid != 0 {
// 		transaksi.Trip_id = tripid
// 	}

// 	data, err := h.TransaksiRepository.UpdateTransaksi(transaksi)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	test, err := h.TransaksiRepository.GetTransaksi(data.ID)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
// 		json.NewEncoder(w).Encode(response)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	response := dto.SuccessResult{Code: http.StatusOK, Data: test}
// 	json.NewEncoder(w).Encode(response)
// }

func (h *handleTransaksi) Notification(w http.ResponseWriter, r *http.Request) {
	var notificationPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&notificationPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)
	id, _ := strconv.Atoi(orderId)

	transaction, _ := h.TransaksiRepository.GetTransaksi(id)
	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {

			h.TransaksiRepository.UpdateTransaksi("pending", transaction.ID)
		} else if fraudStatus == "accept" {
			SendMail("success", transaction)
			h.TransaksiRepository.UpdateTransaksi("success", transaction.ID)
		}
	} else if transactionStatus == "settlement" {
		SendMail("success", transaction)
		h.TransaksiRepository.UpdateTransaksi("success", transaction.ID)
	} else if transactionStatus == "deny" {
		SendMail("failed", transaction)
		h.TransaksiRepository.UpdateTransaksi("failed", transaction.ID)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		SendMail("failed", transaction)
		h.TransaksiRepository.UpdateTransaksi("failed", transaction.ID)
	} else if transactionStatus == "pending" {

		h.TransaksiRepository.UpdateTransaksi("pending", transaction.ID)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handleTransaksi) DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	transaksi, err := h.TransaksiRepository.GetTransaksi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.TransaksiRepository.DeleteTransaksi(transaksi)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)}
	json.NewEncoder(w).Encode(response)

}

func convertResponseTransaksi(u models.Transaction) transaksidto.TransaksiResponse {
	return transaksidto.TransaksiResponse{
		ID: u.ID,
	}
}

func SendMail(status string, transaction models.Transaction) {

	if status != transaction.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = "your email"
		var CONFIG_AUTH_PASSWORD = "your password"

		var tripName = transaction.Trip.Title
		var price = strconv.Itoa(transaction.Trip.Price)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Transaction Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
    <html lang="en">
      <head>
      <meta charset="UTF-8" />
      <meta http-equiv="X-UA-Compatible" content="IE=edge" />
      <meta name="viewport" content="width=device-width, initial-scale=1.0" />
      <title>Document</title>
      <style>
        h1 {
        color: brown;
        }
      </style>
      </head>
      <body>
      <h2>Product payment :</h2>
      <ul style="list-style-type:none;">
        <li>Name : %s</li>
        <li>Total payment: Rp.%s</li>
        <li>Status : <b>%s</b></li>
      </ul>
      </body>
    </html>`, tripName, price, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + transaction.User.Email)
	}
}
