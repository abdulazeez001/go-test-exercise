package models

type VirtualAccountRequest struct {
	Email       string `json:"email"`
	IsPermanent bool   `json:"is_permanent"`
	TXRef       string `json:"tx_ref"`
	Bvn         string `json:"bvn,omitempty"`
}

type VirtualAccountResponse struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	AccountName   string `json:"account_name"`
}
