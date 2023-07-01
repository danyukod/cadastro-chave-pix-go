package services_test

import (
	"errors"
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

func (m *mockRegisterPixKeyRepository) RegisterPixKey(_ pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error) {
	return PixKeyMockFactory()
}

func PixKeyMockFactory() (pix_key.PixKeyDomainInterface, error) {
	pixKeyType := pix_key.PixKeyTypeFromText("cpf")
	accountType := account.AccountTypeFromText("corrente")
	holderDomain, _ := holder.NewHolderDomain("John", "Doe")
	accountDomain, _ := account.NewAccountDomain(123, 1, accountType, holderDomain)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, "39357160876", accountDomain)
	pixKeyDomain.SetID("1")
	return pixKeyDomain, err
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
	request.AccountType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &account.ErrInvalidAccountType{}))
	//
	request.AccountType = "corrente"
	request.PixKeyType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.Is(err, &pix_key.ErrInvalidPixKeyType{}))
	//
	//request.PixKeyType = "email"
	//request.AccountNumber = ""
	//response, err = service.Execute(request)
	//assert.NotNil(t, err)
	//assert.Nil(t, response)
	//assert.True(t, errors.Is(err, account.ErrInvalidAccountNumber))
	//
	//request.AccountNumber = "123456"
	//request.AgencyNumber = ""
	//response, err = service.Execute(request)
	//assert.NotNil(t, err)
	//assert.Nil(t, response)
	//assert.True(t, errors.Is(err, account.ErrInvalidAgencyNumber))
	//
	//request.AgencyNumber = "7890"
	//request.AccountHolderName = ""
	//response, err = service.Execute(request)
	//assert.NotNil(t, err)
	//assert.Nil(t, response)
	//assert.True(t, errors.Is(err, holder.ErrInvalidHolderName))
	//
	//request.AccountHolderName = "John"
	//request.AccountHolderLastName = ""
	//response, err = service.Execute(request)
	//assert.NotNil(t, err)
	//assert.Nil(t, response)
	//assert.True(t, errors.Is(err, holder.ErrInvalidHolderLastName))
}
