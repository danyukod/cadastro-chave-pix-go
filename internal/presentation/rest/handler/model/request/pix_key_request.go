package request

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/application/commands/dto"
)

type RegisterPixKeyRequest struct {
	PixKeyType            string `json:"pixKeyType" binding:"required"`
	PixKey                string `json:"response" binding:"required"`
	AccountType           string `json:"accountType" binding:"required"`
	AccountNumber         int    `json:"accountNumber" binding:"required"`
	AgencyNumber          int    `json:"agencyNumber" binding:"required"`
	AccountHolderName     string `json:"accountHolderName" binding:"required"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}

type FindPixKeyRequest struct {
	PixKeyType string `uri:"pixKeyType" binding:"required"`
	PixKey     string `uri:"response" binding:"required"`
}

func (f *FindPixKeyRequest) ToDTO() dto.FindPixKeyDTO {
	return dto.FindPixKeyDTO{
		Key: f.PixKey,
	}
}

func (f *RegisterPixKeyRequest) ToDTO() dto.RegisterPixKeyDTO {
	return dto.RegisterPixKeyDTO{
		PixKeyType:            f.PixKeyType,
		PixKey:                f.PixKey,
		AccountType:           f.AccountType,
		AccountNumber:         f.AccountNumber,
		AgencyNumber:          f.AgencyNumber,
		AccountHolderName:     f.AccountHolderName,
		AccountHolderLastName: f.AccountHolderLastName,
	}
}
