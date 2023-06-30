package output

import "github.com/danyukod/cadastro-chave-pix-go/src/domain"

type RegisterPixKeyRepository interface {
	RegisterPixKey(
		pixKeyDomain domain.PixKeyDomainInterface,
	) (domain.PixKeyDomainInterface, error)
}
