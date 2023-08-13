package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	shared2 "github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
)

type MockRegisterPixKeyUseCase struct{}

type MockRegisterPixKeyUseCaseError struct{}

type MockFindPixKeyUseCase struct{}

type MockFindPixKeyUseCaseError struct{}

func (m *MockRegisterPixKeyUseCase) Execute(_ requestpackage.RegisterPixKeyRequest) (model.PixKeyDomainInterface, error) {
	holderDomain, _ := shared2.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := shared2.NewAccountDomain(123, 1, shared2.CORRENTE, holderDomain)
	return model.NewPixKeyDomain(shared2.CPF, "39357160876", accountDomain)
}

func (m *MockFindPixKeyUseCase) Execute(_ requestpackage.FindPixKeyRequest) (model.PixKeyDomainInterface, error) {
	holderDomain, _ := shared2.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := shared2.NewAccountDomain(123, 1, shared2.CORRENTE, holderDomain)
	return model.NewPixKeyDomain(shared2.CPF, "39357160876", accountDomain)
}

func (m *MockRegisterPixKeyUseCaseError) Execute(_ requestpackage.RegisterPixKeyRequest) (model.PixKeyDomainInterface, error) {
	var businessErrors shared2.BusinessErrors
	businessErrors = append(businessErrors, *shared2.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}

func (m *MockFindPixKeyUseCaseError) Execute(_ requestpackage.FindPixKeyRequest) (model.PixKeyDomainInterface, error) {
	var businessErrors shared2.BusinessErrors
	businessErrors = append(businessErrors, *shared2.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}
