package mapper

import (
	"github.com/danyukod/cadastro-chave-pix-go/src/adapters/output/database/entity"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/src/domain/pix_key"
	"github.com/google/uuid"
	"time"
)

func ConvertDomainToEntity(domain pix_key.PixKeyDomainInterface) entity.PixKeyEntity {
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

func ConvertEntityToDomain(entity entity.PixKeyEntity) (pix_key.PixKeyDomainInterface, error) {
	holder, err := holder.NewHolderDomain(entity.AccountHolderName, entity.AccountHolderLastName)
	if err != nil {
		return nil, err
	}
	account, err := account.NewAccountDomain(entity.AccountNumber, entity.AgencyNumber, account.AccountTypeFromText(entity.AccountType), holder)
	if err != nil {
		return nil, err
	}
	pixKey, err := pix_key.NewPixKeyDomain(pix_key.PixKeyTypeFromText(entity.PixKeyType), entity.PixKey, account)
	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
