package entity

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/account"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/holder"
	"github.com/danyukod/cadastro-chave-pix-go/internal/app/pix_key/src/domain/pix_key"
	"github.com/google/uuid"
	"time"
)

type PixKeyEntity struct {
	ID                    string    `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	PixKeyType            string    `gorm:"type:varchar(20);not null"`
	PixKey                string    `gorm:"type:varchar(255);not null"`
	AccountType           string    `gorm:"type:varchar(20);not null"`
	AccountNumber         int       `gorm:"type:int;not null"`
	AgencyNumber          int       `gorm:"type:int;not null"`
	AccountHolderName     string    `gorm:"type:varchar(255);not null"`
	AccountHolderLastName string    `gorm:"type:varchar(255);not null"`
	CreatedAt             time.Time `gorm:"type:timestamp;not null"`
	ModifiedAt            time.Time `gorm:"type:timestamp;not null"`
}

func (PixKeyEntity) TableName() string {
	return "pix_key"
}

func ConvertDomainToEntity(domain pix_key.PixKeyDomainInterface) PixKeyEntity {
	var id string
	if domain.GetID() != "" {
		id = domain.GetID()
	} else {
		id = uuid.NewString()
	}
	return PixKeyEntity{
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

func PixKeyDomainFromEntity(entity PixKeyEntity) (pix_key.PixKeyDomainInterface, error) {
	holderDomain, err := holder.NewHolderDomain(entity.AccountHolderName, entity.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountDomain, err := account.NewAccountDomain(entity.AccountNumber, entity.AgencyNumber, account.AccountTypeFromText(entity.AccountType), holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain, err := pix_key.NewPixKeyDomain(pix_key.PixKeyTypeFromText(entity.PixKeyType), entity.PixKey, accountDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain.SetID(entity.ID)

	return pixKeyDomain, nil
}
