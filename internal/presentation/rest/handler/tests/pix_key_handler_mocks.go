package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands/dto"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/value_object"
)

type MockRegisterPixKeyUseCase struct{}

type MockRegisterPixKeyUseCaseError struct{}

type MockFindPixKeyUseCase struct{}

type MockFindPixKeyUseCaseError struct{}

func (m *MockRegisterPixKeyUseCase) Execute(_ dto.RegisterPixKeyDTO) (model.PixKeyDomainInterface, error) {
	holderDomain, _ := aggregate.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE.String(), holderDomain)
	return model.NewPixKeyDomain(value_object.CPF, "39357160876", accountDomain)
}

func (m *MockFindPixKeyUseCase) Execute(_ dto.FindPixKeyDTO) (model.PixKeyDomainInterface, error) {
	holderDomain, _ := aggregate.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE.String(), holderDomain)
	return model.NewPixKeyDomain(value_object.CPF, "39357160876", accountDomain)
}

func (m *MockRegisterPixKeyUseCaseError) Execute(_ dto.RegisterPixKeyDTO) (model.PixKeyDomainInterface, error) {
	var businessErrors value_object.BusinessErrors
	businessErrors = append(businessErrors, *value_object.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}

func (m *MockFindPixKeyUseCaseError) Execute(_ dto.FindPixKeyDTO) (model.PixKeyDomainInterface, error) {
	var businessErrors value_object.BusinessErrors
	businessErrors = append(businessErrors, *value_object.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}
