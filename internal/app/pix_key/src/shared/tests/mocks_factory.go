package tests

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
)

func PixKeyMockFactory() (pix_key.PixKeyDomainInterface, error) {
	pixKeyType := pix_key.PixKeyTypeFromText("cpf")
	accountType := account.AccountTypeFromText("corrente")
	holderDomain, _ := holder.NewHolderDomain("John", "Doe")
	accountDomain, _ := account.NewAccountDomain(123, 1, accountType, holderDomain)
	pixKeyDomain, err := pix_key.NewPixKeyDomain(pixKeyType, "39357160876", accountDomain)
	return pixKeyDomain, err
}
