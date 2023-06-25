package response

type RegisterPixKeyResponse struct {
	Id                    string `json:"id"`
	PixKeyType            string `json:"pixKeyType"`
	PixKey                string `json:"pixKey"`
	AccountType           string `json:"accountType"`
	AccountNumber         string `json:"accountNumber"`
	AgencyNumber          string `json:"agencyNumber"`
	AccountHolderName     string `json:"accountHolderName"`
	AccountHolderLastName string `json:"accountHolderLastName"`
	CreatedAt             string `json:"createdAt"`
	UpdatedAt             string `json:"updatedAt"`
}
