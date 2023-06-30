package mapper

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/enum"
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
		AgencyNumber:          domain.GetAccount().GetAgency(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
		CreatedAt:             time.Now(),
		ModifiedAt:            time.Now(),
	}
}

func ConvertEntityToDomain(entity entity.PixKeyEntity) (domain.PixKeyDomainInterface, error) {
	holder, err := domain.NewHolderDomain(entity.AccountHolderName, entity.AccountHolderLastName)
	if err != nil {
		return nil, err
	}
	account, err := domain.NewAccountDomain(entity.AccountNumber, entity.AgencyNumber, enum.AccountTypeFromText(entity.AccountType), holder)
	if err != nil {
		return nil, err
	}
	pixKey, err := domain.NewPixKeyDomain(enum.PixKeyTypeFromText(entity.PixKeyType), entity.PixKey, account)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
