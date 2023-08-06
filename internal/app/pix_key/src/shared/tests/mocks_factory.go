package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/input/web/controller/model/request"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
)

func PixKeyMockFactory() (pix_key.PixKeyDomainInterface, error) {
	pixKeyRequest := request.RegisterPixKeyRequest{
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
	}
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyRequest)
	return pixKeyDomain, err
}
