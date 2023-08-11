package input_test

import (
	"errors"
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/model/response"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/shared/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRegisterPixKeyRepository struct{}

type mockRegisterPixKeyRepositoryWithError struct{}

func (m *mockRegisterPixKeyRepository) RegisterPixKey(_ domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	pixKeyDomain, err := tests.PixKeyMockFactory()
	if err != nil {
		return nil, err
	}
	pixKeyDomain.SetID("1")
	return pixKeyDomain, nil
}

func (m *mockRegisterPixKeyRepository) VerifyIfPixKeyAlreadyExists(_ string, _ string) error {
	return nil
}

func (m *mockRegisterPixKeyRepositoryWithError) RegisterPixKey(_ domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	return nil, businesserrors.BusinessErrors{*businesserrors.NewBusinessError("PixKey", "PixKey already exists", "123")}
}

func (m *mockRegisterPixKeyRepositoryWithError) VerifyIfPixKeyAlreadyExists(_ string, _ string) error {
	return businesserrors.BusinessErrors{*businesserrors.NewBusinessError("PixKey", "PixKey already exists", "123")}
}

func TestRegisterPixKeyService_Execute(t *testing.T) {
	pixKeyRepository := &mockRegisterPixKeyRepository{}
	service := NewRegisterPixKeyService(pixKeyRepository)

	request := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}

	expectedResponse := response.FindPixKeyResponse{
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

	var businessErrors businesserrors.BusinessErrors

	// Test handler handling
	// Invalid Account Type
	request.AccountType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Type", businessErrors[0].Field())
	// Invalid Account Number
	request.AccountType = "corrente"
	request.AccountNumber = 0
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Number", businessErrors[0].Field())
	// Invalid Account Agency
	request.AccountNumber = 1
	request.AgencyNumber = 0
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da agencia esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Agency Number", businessErrors[0].Field())
	// Invalid Holder Name
	request.AgencyNumber = 1
	request.AccountHolderName = ""
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O nome do titular esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Holder Name", businessErrors[0].Field())
	//Invalid Pix Key Type
	request.AccountHolderName = "Joe"
	request.PixKeyType = "invalid"
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de chave esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key Type", businessErrors[0].Field())
	//Invalid Pix Key
	request.PixKeyType = "cpf"
	request.PixKey = ""
	response, err = service.Execute(request)
	assert.NotNil(t, err)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O valor da chave esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key", businessErrors[0].Field())

}

func TestRegisterPixKeyService_ExecuteWithError(t *testing.T) {
	pixKeyRepository := &mockRegisterPixKeyRepositoryWithError{}
	service := NewRegisterPixKeyService(pixKeyRepository)

	request := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}

	// Test handler handling
	response, err := service.Execute(request)
	var businessErrors businesserrors.BusinessErrors
	errors.As(err, &businessErrors)

	assert.NotNil(t, businessErrors)
	assert.Nil(t, response)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "PixKey already exists", businessErrors[0].Error())
	assert.Equal(t, "PixKey", businessErrors[0].Field())
}
