package commands

import (
	"errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type PixKeyPersistenceInterfaceMock struct {
	mock.Mock
}

func (m *PixKeyPersistenceInterfaceMock) FindPixKey() ([]model.PixKeyDomainInterface, error) {
	args := m.Called()
	return args.Get(0).([]model.PixKeyDomainInterface), args.Error(1)
}

func (m *PixKeyPersistenceInterfaceMock) FindById(id string) (model.PixKeyDomainInterface, error) {
	args := m.Called(id)
	return args.Get(0).(model.PixKeyDomainInterface), args.Error(1)
}

func (m *PixKeyPersistenceInterfaceMock) CreatePixKey(pixKeyDomain model.PixKeyDomainInterface) (model.PixKeyDomainInterface, error) {
	args := m.Called(pixKeyDomain)
	return args.Get(0).(model.PixKeyDomainInterface), args.Error(1)
}

func (m *PixKeyPersistenceInterfaceMock) FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (model.PixKeyDomainInterface, error) {
	args := m.Called(pixKeyType, pixKey)
	return args.Get(0).(model.PixKeyDomainInterface), args.Error(1)
}

func TestFindPixKeyService_Execute(t *testing.T) {
	pm := PixKeyPersistenceInterfaceMock{}
	cpfId := "39357160876"

	holderDomain, err := aggregate.NewHolderDomain("John", "Doe")
	assert.Nil(t, err)
	assert.NotNil(t, holderDomain)
	accountDomain, err := aggregate.NewAccountDomain(123, 1, "CORRENTE", holderDomain)
	assert.Nil(t, err)
	assert.NotNil(t, accountDomain)
	pixKeyDomain, err := model.NewPixKeyDomain(value_object.CPF, cpfId, accountDomain)
	assert.Nil(t, err)
	assert.NotNil(t, pixKeyDomain)

	pm.On("FindById", cpfId).Return(pixKeyDomain, nil)

	service := NewFindPixKeyService(&pm)

	pixKeyDomainList, err := service.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, pixKeyDomainList)
	assert.Equal(t, 1, len(pixKeyDomainList))
	assert.Equal(t, cpfId, pixKeyDomainList[0].GetPixKey())

}

func TestFindPixKeyService_Execute_Error(t *testing.T) {
	pm := PixKeyPersistenceInterfaceMock{}
	cpfId := "39357160876"
	domain, _ := model.NewPixKeyDomain(value_object.CPF, cpfId, nil)
	pm.On("FindById", cpfId).Return(domain, errors.New("Pix Key not found"))

	service := NewFindPixKeyService(&pm)

	pixKeyDomain, err := service.Execute()

	assert.Nil(t, pixKeyDomain)
	assert.NotNil(t, err)
	assert.Equal(t, "Pix Key not found", err.Error())
}
