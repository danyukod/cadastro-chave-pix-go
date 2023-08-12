package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/ui/adapter/rest/controller/model/request"
)

func PixKeyMockFactory() (domain.PixKeyDomainInterface, error) {
	pixKeyRequest := request.RegisterPixKeyRequest{
		PixKeyType:            "cpf",
		PixKey:                "39357160876",
		AccountType:           "corrente",
		AccountNumber:         123,
		AgencyNumber:          1,
		AccountHolderName:     "John",
		AccountHolderLastName: "Doe",
	}
	pixKeyDomain, err := domain.PixKeyDomainFromRequest(pixKeyRequest)
	return pixKeyDomain, err
}
