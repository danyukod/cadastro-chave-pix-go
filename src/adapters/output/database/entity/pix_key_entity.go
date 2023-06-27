package entity

import "time"

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
