package mapper

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/adapters/output/database/entity"
	account3 "github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
	"github.com/google/uuid"
	"time"
)

func ConvertDomainToEntity(domain pix_key.PixKeyDomainInterface) entity.PixKeyEntity {
	var id string
	if domain.GetID() != "" {
		id = domain.GetID()
	} else {
		id = uuid.NewString()
	}
	return entity.PixKeyEntity{
		ID:                    id,
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
	account, err := account3.NewAccountDomain(entity.AccountNumber, entity.AgencyNumber, account3.AccountTypeFromText(entity.AccountType), holder)
	if err != nil {
		return nil, err
	}
	pixKey, err := pix_key.NewPixKeyDomain(pix_key.PixKeyTypeFromText(entity.PixKeyType), entity.PixKey, account)
	if err != nil {
		return nil, err
	}

	pixKey.SetID(entity.ID)

	return pixKey, nil
}
