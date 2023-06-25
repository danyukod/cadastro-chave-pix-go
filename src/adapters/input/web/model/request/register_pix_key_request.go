package request

type RegisterPixKeyRequest struct {
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"pixKey"`
	AccountType           string `json:"accountType"`
	AccountNumber         string `json:"accountNumber"`
	AgencyNumber          string `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
}
