package dto

type RegisterPixKeyDTO struct {
	PixKeyType            string
	PixKey                string
	AccountType           string
	AccountNumber         int
	AgencyNumber          int
	AccountHolderName     string
	AccountHolderLastName string
}

type FindPixKeyDTO struct {
	Key string `uri:"key" binding:"required"`
}
