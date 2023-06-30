package output

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
)

type RegisterPixKeyRepository interface {
	RegisterPixKey(
		pixKeyDomain pix_key.PixKeyDomainInterface,
	) (pix_key.PixKeyDomainInterface, error)
}
