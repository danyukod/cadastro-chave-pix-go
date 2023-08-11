package tests

import (
	businesserrors "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/application/errors"
	requestpackage "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/infrastructure/adapters/input/web/controller/model/response"
)

type MockRegisterPixKeyUseCase struct{}

type MockRegisterPixKeyUseCaseError struct{}

type MockFindPixKeyUseCase struct{}

type MockFindPixKeyUseCaseError struct{}

func (m *MockRegisterPixKeyUseCase) Execute(_ requestpackage.RegisterPixKeyRequest) (*response.FindPixKeyResponse, error) {
	return &response.FindPixKeyResponse{
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

func (m *MockFindPixKeyUseCase) Execute(_ requestpackage.FindPixKeyRequest) (*response.FindPixKeyResponse, error) {
	return &response.FindPixKeyResponse{
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

func (m *MockRegisterPixKeyUseCaseError) Execute(_ requestpackage.RegisterPixKeyRequest) (*response.FindPixKeyResponse, error) {
	var businessErrors businesserrors.BusinessErrors
	businessErrors = append(businessErrors, *businesserrors.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}

func (m *MockFindPixKeyUseCaseError) Execute(_ requestpackage.FindPixKeyRequest) (*response.FindPixKeyResponse, error) {
	var businessErrors businesserrors.BusinessErrors
	businessErrors = append(businessErrors, *businesserrors.NewBusinessError("Pix Key", "O valor da chave esta invalido.", "123"))
	return nil, businessErrors
}
