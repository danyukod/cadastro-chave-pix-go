package output

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
)

type PixKeyPersistenceInterface interface {
	CreatePixKey(pix_key.PixKeyDomainInterface) (pix_key.PixKeyDomainInterface, error)
	FindPixKeyByKeyAndType(pixKeyType string, pixKey string) (pix_key.PixKeyDomainInterface, error)
}