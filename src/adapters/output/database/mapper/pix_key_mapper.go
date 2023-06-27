package mapper

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
	"github.com/google/uuid"
	"time"
)

func ConvertDomainToEntity(domain domain.PixKeyDomainInterface) entity.PixKeyEntity {
	return entity.PixKeyEntity{
		ID:                    uuid.NewString(),
		PixKeyType:            domain.GetPixKeyType().String(),
		PixKey:                domain.GetPixKey(),
		AccountType:           domain.GetAccount().GetAccountType().String(),
		AccountNumber:         domain.GetAccount().GetNumber(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}
}
