package output

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
)

type PixKeyPersistenceInterface interface {
	CreatePixKey(pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error)
	FindPixKeyByType(pixKeyType string) (pix_key.PixKeyDomainInterface, error)
}
