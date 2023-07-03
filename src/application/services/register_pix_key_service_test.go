package services_test

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/src/shared/tests"
	"testing"

	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/src/application/services"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
	"github.com/stretchr/testify/assert"
)

type mockRegisterPixKeyRepository struct{}

type mockRegisterPixKeyRepositoryWithError struct{}

func (m *mockRegisterPixKeyRepository) RegisterPixKey(_ pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
	return tests.PixKeyMockFactory()
}

func (m *mockRegisterPixKeyRepository) VerifyIfPixKeyAlreadyExists(_ string) (bool, error) {
	return false, nil
}

func (m *mockRegisterPixKeyRepositoryWithError) RegisterPixKey(_ pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
	return nil, &pix_key.ErrPixKeyAlreadyExists{}
}

func (m *mockRegisterPixKeyRepositoryWithError) VerifyIfPixKeyAlreadyExists(_ string) (bool, error) {
	return true, nil
}

func TestRegisterPixKeyService_Execute(t *testing.T) {
	pixKeyRepository := &mockRegisterPixKeyRepository{}
	service := services.NewRegisterPixKeyService(pixKeyRepository)

	request := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}

	expectedResponse := response.RegisterPixKeyResponse{
		Id:                    "1",
		PixKeyType:            "CPF",
		PixKey:                "39357160876",
		AccountType:           "CORRENTE",
		AccountNumber:         123,
		AgencyNumber:          1,
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
	}

	// Test successful execution
	response, err := service.Execute(request)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, *response)

	// Test error handling
	// Invalid Account Type
	request.AccountType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &account.ErrInvalidAccountType{}))
	// Invalid Account Number
	request.AccountType = "corrente"
	request.AccountNumber = 0
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &account.ErrInvalidAccountNumber{}))
	// Invalid Account Agency
	request.AccountNumber = 1
	request.AgencyNumber = 0
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &account.ErrInvalidAccountAgency{}))
	// Invalid Holder Name
	request.AgencyNumber = 1
	request.AccountHolderName = ""
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &holder.ErrInvalidHolderName{}))
	//Invalid Pix Key Type
	request.AccountHolderName = "Joe"
	request.PixKeyType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &pix_key.ErrInvalidPixKeyType{}))
	//Invalid Pix Key
	request.PixKeyType = "cpf"
	request.PixKey = ""
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &pix_key.ErrInvalidPixKey{}))

}

func TestRegisterPixKeyService_ExecuteWithError(t *testing.T) {
	pixKeyRepository := &mockRegisterPixKeyRepositoryWithError{}
	service := services.NewRegisterPixKeyService(pixKeyRepository)

	request := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}

	// Test error handling
	response, err := service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &pix_key.ErrPixKeyAlreadyExists{}))
}
