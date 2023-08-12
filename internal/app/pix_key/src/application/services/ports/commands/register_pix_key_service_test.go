package commands_test

import (
	"errors"
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/services/ports/commands"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/object_value"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapter/orm"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPixKeyPersistence struct {
	persistence orm.PixKeyPersistenceInterface
}

func (m mockPixKeyPersistence) CreatePixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	return pixKeyDomain, nil
}

func (m mockPixKeyPersistence) FindPixKeyByKeyAndType(_ string, _ string) (domain.PixKeyDomainInterface, error) {
	return nil, nil
}

func (m mockPixKeyPersistence) FindById(id string) (domain.PixKeyDomainInterface, error) {
	holder, _ := aggregate.NewHolderDomain("John", "Doe")
	account, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE, holder)
	return domain.NewPixKeyDomain(object_value.CPF, "39357160876", account)
}

type mockPixKeyPersistenceWithError struct {
	persistence orm.PixKeyPersistenceInterface
}

func (m mockPixKeyPersistenceWithError) CreatePixKey(pixKeyDomain domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserrors.BusinessErrors
	return nil, businesserrors.AddError(businessErrors, *businesserrors.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "pixKey"))
}

func (m mockPixKeyPersistenceWithError) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserrors.BusinessErrors
	return nil, businesserrors.AddError(businessErrors, *businesserrors.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "pixKey"))

}

func (m mockPixKeyPersistenceWithError) FindById(id string) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserrors.BusinessErrors
	return nil, businesserrors.AddError(businessErrors, *businesserrors.NewBusinessError("Pix Key", "Chave pix ja cadastrada.", "pixKey"))

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
	account, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE, holder)

	pixKeyDomainResponse, _ := domain.NewPixKeyDomain(object_value.CPF, "39357160876", account)

	// Test successful execution
	pixKeyDomain, err := service.Execute(pixKeyRequest)
	assert.Nil(t, err)
	assert.Equal(t, pixKeyDomainResponse, pixKeyDomain)

	var businessErrors businesserrors.BusinessErrors

	// Test handler handling
	// Invalid Account Type
	pixKeyRequest.AccountType = "invalid"
	pixKeyDomain, err = service.Execute(pixKeyRequest)
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Type", businessErrors[0].Field())
	// Invalid Account Number
	pixKeyRequest.AccountType = "corrente"
	pixKeyRequest.AccountNumber = 0
	pixKeyDomain, err = service.Execute(pixKeyRequest)
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da conta esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Number", businessErrors[0].Field())
	// Invalid Account Agency
	pixKeyRequest.AccountNumber = 1
	pixKeyRequest.AgencyNumber = 0
	pixKeyDomain, err = service.Execute(pixKeyRequest)
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O numero da agencia esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Agency Number", businessErrors[0].Field())
	// Invalid Holder Name
	pixKeyRequest.AgencyNumber = 1
	pixKeyRequest.AccountHolderName = ""
	pixKeyDomain, err = service.Execute(pixKeyRequest)
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O nome do titular esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Account Holder Name", businessErrors[0].Field())
	//Invalid Pix Key Type
	pixKeyRequest.AccountHolderName = "Joe"
	pixKeyRequest.PixKeyType = "invalid"
	pixKeyDomain, err = service.Execute(pixKeyRequest)
	assert.NotNil(t, err)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "O tipo de chave esta invalido.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key Type", businessErrors[0].Field())
	//Invalid Pix Key
	pixKeyRequest.PixKeyType = "cpf"
	pixKeyRequest.PixKey = ""
	pixKeyDomain, err = service.Execute(pixKeyRequest)
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
	pixKeyDomain, err := service.Execute(pixKeyRequest)
	var businessErrors businesserrors.BusinessErrors
	errors.As(err, &businessErrors)

	assert.NotNil(t, businessErrors)
	assert.Nil(t, pixKeyDomain)
	assert.True(t, errors.As(err, &businessErrors))
	assert.Equal(t, "Chave pix ja cadastrada.", businessErrors[0].Error())
	assert.Equal(t, "Pix Key", businessErrors[0].Field())
}
