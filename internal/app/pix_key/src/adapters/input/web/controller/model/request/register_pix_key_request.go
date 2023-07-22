package request

type RegisterPixKeyRequest struct {
	PixKeyType            string `json:"pixKeyType" binding:"required"`
	PixKey                string `json:"pixKey" binding:"required"`
	AccountType           string `json:"accountType" binding:"required"`
	AccountNumber         int    `json:"accountNumber" binding:"required"`
	AgencyNumber          int    `json:"agencyNumber" binding:"required"`
	AccountHolderName     string `json:"accountHolderName" binding:"required"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}
