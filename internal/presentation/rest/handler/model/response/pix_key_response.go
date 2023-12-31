package response

import (
	"github.com/danyukod/cadastro-chave-pix-go/internal/domain/model"
)

type FindPixKeyResponse struct {
	Id                    string `json:"id"`
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}

func PixKeyDomainToFindWebResponse(domain model.PixKeyDomainInterface) *FindPixKeyResponse {
	return &FindPixKeyResponse{
		Id:                    domain.GetID(),
		PixKeyType:            domain.GetPixKeyType().GetType(),
		PixKey:                domain.GetPixKey(),
		AccountType:           domain.GetAccount().GetAccountType().String(),
		AccountNumber:         domain.GetAccount().GetNumber(),
		AgencyNumber:          domain.GetAccount().GetAgency(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
	}
}

type RegisterPixKeyResponse struct {
	Id                    string `json:"id"`
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"response"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}

func PixKeyDomainToRegisterWebResponse(domain model.PixKeyDomainInterface) *RegisterPixKeyResponse {
	return &RegisterPixKeyResponse{
		Id:                    domain.GetID(),
		PixKeyType:            domain.GetPixKeyType().GetType(),
		PixKey:                domain.GetPixKey(),
		AccountType:           domain.GetAccount().GetAccountType().String(),
		AccountNumber:         domain.GetAccount().GetNumber(),
		AgencyNumber:          domain.GetAccount().GetAgency(),
		AccountHolderName:     domain.GetAccount().GetHolder().GetName(),
		AccountHolderLastName: domain.GetAccount().GetHolder().GetLastName(),
	}
}
