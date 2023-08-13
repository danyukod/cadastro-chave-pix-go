package tests

import (
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/value_object"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/handler/model/request"
)

type MockRegisterPixKeyUseCase struct{}

type MockRegisterPixKeyUseCaseError struct{}

type MockFindPixKeyUseCase struct{}

type MockFindPixKeyUseCaseError struct{}

func (m *MockRegisterPixKeyUseCase) Execute(_ requestpackage.RegisterPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	holderDomain, _ := aggregate.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE, holderDomain)
	return domain.NewPixKeyDomain(value_object.CPF, "39357160876", accountDomain)
}

func (m *MockFindPixKeyUseCase) Execute(_ requestpackage.FindPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	holderDomain, _ := aggregate.NewHolderDomain("Danilo", "Kodavara")
	accountDomain, _ := aggregate.NewAccountDomain(123, 1, aggregate.CORRENTE, holderDomain)
	return domain.NewPixKeyDomain(value_object.CPF, "39357160876", accountDomain)
}

func (m *MockRegisterPixKeyUseCaseError) Execute(_ requestpackage.RegisterPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserrors.BusinessErrors
	businessErrors = append(businessErrors, *businesserrors.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}

func (m *MockFindPixKeyUseCaseError) Execute(_ requestpackage.FindPixKeyRequest) (domain.PixKeyDomainInterface, error) {
	var businessErrors businesserrors.BusinessErrors
	businessErrors = append(businessErrors, *businesserrors.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}
