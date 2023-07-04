package controller

import (
	"bytes"
	"encoding/json"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/request"
	response2 "github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/response"
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/src/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRegisterPixKeyUseCase struct{}

type mockRegisterPixKeyUseCaseError struct{}

func (m *mockRegisterPixKeyUseCase) Execute(_ requestpackage.RegisterPixKeyRequest) (*response2.RegisterPixKeyResponse, error) {
	return &response2.RegisterPixKeyResponse{
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

func (m *mockRegisterPixKeyUseCaseError) Execute(_ requestpackage.RegisterPixKeyRequest) (*response2.RegisterPixKeyResponse, error) {
	var businessErrors businesserrors.BusinessErrors
	businessErrors = append(businessErrors, *businesserrors.NewBusinessError("Pix Key", "O valor da chave esta invalido."))
	return nil, businessErrors
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

		var response response2.RegisterPixKeyResponse
		json.Unmarshal(w.Body.Bytes(), &response)

		assert.NotNil(t, response)
		assert.Equal(t, "CPF", response.PixKeyType)
		assert.Equal(t, "39357160876", response.PixKey)
		assert.Equal(t, "CORRENTE", response.AccountType)
		assert.Equal(t, 123, response.AccountNumber)
		assert.Equal(t, 1, response.AgencyNumber)
		assert.Equal(t, "Danilo", response.AccountHolderName)
		assert.Equal(t, "Kodavara", response.AccountHolderLastName)
	})

	t.Run("should return 400 status code and error message when invalid PixKey", func(t *testing.T) {

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

		var response map[string][]response2.ErrorResponse

		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "Pix Key", response["errors"][0].Field)
		assert.Equal(t, "O valor da chave esta invalido.", response["errors"][0].Message)
	})

	t.Run("should return 400 status code when invalid request body", func(t *testing.T) {

		mockUseCase := &mockRegisterPixKeyUseCaseError{}
		controller := NewPixKeyControllerInterface(mockUseCase)

		router := gin.Default()
		router.POST("/pix-key", controller.RegisterPixKey)

		requestBody, _ := json.Marshal(requestpackage.RegisterPixKeyRequest{
			PixKeyType:            "",
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

		var response map[string][]response2.ErrorResponse

		json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, "PixKeyType", response["errors"][0].Field)
		assert.Equal(t, "This field is required", response["errors"][0].Message)

	})
}
