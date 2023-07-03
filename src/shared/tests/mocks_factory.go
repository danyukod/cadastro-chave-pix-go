package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
)

func PixKeyMockFactory() (pix_key.PixKeyDomainInterface, error) {
	pixKeyType := pix_key.PixKeyTypeFromText("cpf")
	accountType := account.AccountTypeFromText("corrente")
	holderDomain, _ := holder.NewHolderDomain("John", "Doe")
	accountDomain, _ := account.NewAccountDomain(123, 1, accountType, holderDomain)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, "39357160876", accountDomain)
	pixKeyDomain.SetID("1")
	return pixKeyDomain, err
}
