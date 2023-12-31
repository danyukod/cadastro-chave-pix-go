package entity

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/shared/aggregate"
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

func ConvertDomainToEntity(domain model.PixKeyDomainInterface) PixKeyEntity {
	var id string
	if domain.GetID() != "" {
		id = domain.GetID()
	} else {
		id = uuid.NewString()
	}
	return PixKeyEntity{
		ID:                    id,
		PixKeyType:            domain.GetPixKeyType().GetType(),
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

func PixKeyDomainFromEntity(entity PixKeyEntity) (model.PixKeyDomainInterface, error) {
	holderDomain, err := aggregate.NewHolderDomain(entity.AccountHolderName, entity.AccountHolderLastName)
	if err != nil {
		return nil, err
	}

	accountDomain, err := aggregate.NewAccountDomain(entity.AccountNumber, entity.AgencyNumber, entity.AccountType, holderDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain, err := model.NewPixKeyDomain(entity.PixKeyType, entity.PixKey, accountDomain)
	if err != nil {
		return nil, err
	}

	pixKeyDomain.SetID(entity.ID)

	return pixKeyDomain, nil
}
