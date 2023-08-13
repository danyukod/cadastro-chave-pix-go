package commands_test

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
	"github.com/danyukod/cadastro-chave-pix-go/internal/infrastructure/persistence"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/rest/handler/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPixKeyPersistence struct {
	persistence persistence.PixKeyPersistenceInterface
}

func (m mockPixKeyPersistence) CreatePixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
	return pixKeyDomain, nil
}

func (m mockPixKeyPersistence) FindPixKeyByKeyAndType(_ string, _ string) (model.PixKeyDomainInterface, error) {
	return nil, nil
}

func (m mockPixKeyPersistence) FindById(id string) (model.PixKeyDomainInterface, error) {
	holder, _ := aggregate.NewHolderDomain("John", "Doe")
	account, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE.String(), holder)
	return model.NewPixKeyDomain(value_object.CPF.String(), "39357160876", account)
}

type mockPixKeyPersistenceWithError struct {
	persistence persistence.PixKeyPersistenceInterface
}

func (m mockPixKeyPersistenceWithError) CreatePixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
	var businessErrors value_object.BusinessErrors
	return nil, value_object.AddError(businessErrors, *value_object.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "response"))
}

func (m mockPixKeyPersistenceWithError) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (model.PixKeyDomainInterface, error) {
	var businessErrors value_object.BusinessErrors
	return nil, value_object.AddError(businessErrors, *value_object.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "response"))

}

func (m mockPixKeyPersistenceWithError) FindById(id string) (model.PixKeyDomainInterface, error) {
	var businessErrors value_object.BusinessErrors
	return nil, value_object.AddError(businessErrors, *value_object.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "response"))

}
func TestRegisterPixKeyService_Execute(t *testing.T) {
	persistence := &mockPixKeyPersistence{}
	service := commands.NewRegisterPixKeyService(persistence)

	pixKeyRequest := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}
	holder, _ := aggregate.NewHolderDomain("John", "Doe")
	account, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE.String(), holder)

	pixKeyDomainResponse, _ := model.NewPixKeyDomain(value_object.CPF.String(), "39357160876", account)

	// Test successful execution
	pixKeyDomain, err := service.Execute(pixKeyRequest.ToDTO())
	assert.Nil(t, err)
	assert.Equal(t, pixKeyDomainResponse, pixKeyDomain)

	var businessErrors value_object.BusinessErrors

	// Test handler handling
	// Invalid Account Type
	pixKeyRequest.AccountType = "invalid"
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Type", businessErrors[0].Field())
	// Invalid Account Number
	pixKeyRequest.AccountType = "corrente"
	pixKeyRequest.AccountNumber = 0
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Number", businessErrors[0].Field())
	// Invalid Account Agency
	pixKeyRequest.AccountNumber = 1
	pixKeyRequest.AgencyNumber = 0
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da agencia esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Agency Number", businessErrors[0].Field())
	// Invalid Holder Name
	pixKeyRequest.AgencyNumber = 1
	pixKeyRequest.AccountHolderName = ""
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O nome do titular esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Holder Name", businessErrors[0].Field())
	//Invalid Pix Key Type
	pixKeyRequest.AccountHolderName = "Joe"
	pixKeyRequest.PixKeyType = "invalid"
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de chave esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key Type", businessErrors[0].Field())
	//Invalid Pix Key
	pixKeyRequest.PixKeyType = "cpf"
	pixKeyRequest.PixKey = ""
	pixKeyDomain, err = service.Execute(pixKeyRequest.ToDTO())
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O valor da chave esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key", businessErrors[0].Field())

}

func TestRegisterPixKeyService_ExecuteWithError(t *testing.T) {
	persistence := &mockPixKeyPersistenceWithError{}
	service := commands.NewRegisterPixKeyService(persistence)

	pixKeyRequest := request.RegisterPixKeyRequest{
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
	}

	// Test handler handling
	pixKeyDomain, err := service.Execute(pixKeyRequest.ToDTO())
	var businessErrors value_object.BusinessErrors
	errors.As(err, &businessErrors)

	assert.NotNil(t, businessErrors)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "Chave pix ja cadastrada.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key", businessErrors[0].Field())
}
