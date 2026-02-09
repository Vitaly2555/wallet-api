package handler

type OperationRequest struct {
	WalletID  string `json:"walletId"`
	Operation string `json:"operationType"`
	Amount    int64  `json:"amount"`
}
