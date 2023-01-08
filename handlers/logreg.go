package handlers

import (
	registerdto "Backend/dto/register"
	dto "Backend/dto/result"
	"Backend/models"
	"Backend/pkg/bcrypt"
	jwtToken "Backend/pkg/jwt"
	"fmt"
	"log"
	"time"

	"Backend/repositories"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
)

type handlerLogReg struct {
	LogRegRepository repositories.LogRegRepository
}

func HandlerLogReg(LogRegRepository repositories.LogRegRepository) *handlerLogReg {
	return &handlerLogReg{LogRegRepository}
}

func (h *handlerLogReg) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(registerdto.RegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user := models.User{
		Fullname: request.Fullname,
		Email:    request.Email,
		Password: password,
		Gender:   request.Gender,
		Phone:    request.Phone,
		Address:  request.Address,
		Role:     request.Role,
	}

	data, err := h.LogRegRepository.Register(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseRegister(data)}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerLogReg) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(registerdto.LoginResponse)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		ID:       request.ID,
		Email:    request.Email,
		Password: request.Password,
		Role:			request.Role,
	}

	user, err := h.LogRegRepository.Login(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		fmt.Println("Unauthorize")
		return
	}

	loginResponse := registerdto.LoginResponse{
		ID:       user.ID,
		Fullname: user.Fullname,
		Email:    user.Email,
		Password: user.Password,
		Token:    token,
		Role: user.Role,
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(w).Encode(response)

}

func convertResponseRegister(u models.User) registerdto.RegisterRespons {
	return registerdto.RegisterRespons{
		Email:    u.Email,
		Password: u.Password,
	}
}
