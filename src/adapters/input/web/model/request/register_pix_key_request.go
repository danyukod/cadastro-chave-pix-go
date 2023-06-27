package request

type RegisterPixKeyRequest struct {
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"pixKey"`
	AccountType           string `json:"accountType"`
	AccountNumber         int    `json:"accountNumber"`
	AgencyNumber          int    `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}
