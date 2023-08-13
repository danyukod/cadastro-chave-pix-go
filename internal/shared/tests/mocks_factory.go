package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
	"github.com/danyukod/cadastro-chave-pix-go/internal/presentation/handler/model/request"
)

func PixKeyMockFactory() (model.PixKeyDomainInterface, error) {
	pixKeyRequest := request.RegisterPixKeyRequest{
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
	}
	holderDomain, err := aggregate.NewHolderDomain(pixKeyRequest.AccountHolderName, pixKeyRequest.AccountHolderLastName)

	accoutDomain, err := aggregate.NewAccountDomain(pixKeyRequest.AccountNumber, pixKeyRequest.AgencyNumber, pixKeyRequest.AccountType, holderDomain)

	pixKeyDomain, err := model.NewPixKeyDomain(pixKeyRequest.PixKeyType, pixKeyRequest.PixKey, accoutDomain)

	return pixKeyDomain, err
}
