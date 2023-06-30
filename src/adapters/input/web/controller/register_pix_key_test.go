package controller

import (
	"bytes"
	"encoding/json"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/model/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRegisterPixKeyUseCase struct{}

type mockRegisterPixKeyUseCaseError struct{}

func (m *mockRegisterPixKeyUseCase) Execute(_ requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {
	return &response.RegisterPixKeyResponse{
		Id:                    "12345678900",
		PixKeyType:            "CPF",
		PixKey:                "39357160876",
		AccountType:           "CORRENTE",
		AccountNumber:         123,
		AgencyNumber:          1,
		AccountHolderName:     "Danilo",
		AccountHolderLastName: "Kodavara",
	}, nil
}

func (m *mockRegisterPixKeyUseCaseError) Execute(_ requestpackage.RegisterPixKeyRequest) (*response.RegisterPixKeyResponse, error) {
	return nil, assert.AnError
}

func TestPixKeyController_RegisterPixKey(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("should return 201 status code and success message", func(t *testing.T) {

		mockUseCase := &mockRegisterPixKeyUseCase{}
		controller := NewPixKeyControllerInterface(mockUseCase)

		router := gin.Default()
		router.POST("/pix-key", controller.RegisterPixKey)

		requestBody, _ := json.Marshal(requestpackage.RegisterPixKeyRequest{
			PixKeyType:            "CPF",
			PixKey:                "39357160876",
			AccountType:           "CORRENTE",
			AccountNumber:         123,
			AgencyNumber:          1,
			AccountHolderName:     "Danilo",
			AccountHolderLastName: "Kodavara",
		})

		req, _ := http.NewRequest("POST", "/pix-key", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
	})

	t.Run("should return 400 status code and error message", func(t *testing.T) {

		mockUseCase := &mockRegisterPixKeyUseCaseError{}
		controller := NewPixKeyControllerInterface(mockUseCase)

		router := gin.Default()
		router.POST("/pix-key", controller.RegisterPixKey)

		requestBody, _ := json.Marshal(requestpackage.RegisterPixKeyRequest{
			PixKeyType:            "CPF",
			PixKey:                "39357160876",
			AccountType:           "CORRENTE",
			AccountNumber:         123,
			AgencyNumber:          1,
			AccountHolderName:     "Danilo",
			AccountHolderLastName: "Kodavara",
		})
		req, _ := http.NewRequest("POST", "/pix-key", bytes.NewBuffer(requestBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		json.Unmarshal(w.Body.Bytes(), &response)
	})
}
