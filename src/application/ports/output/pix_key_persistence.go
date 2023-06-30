package output

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
)

type PixKeyPersistenceInterface interface {
	CreatePixKey(domain.PixKeyDomainInterface) (domain.PixKeyDomainInterface, error)
}
