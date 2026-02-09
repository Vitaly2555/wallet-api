package domain

import "errors"

var (
	ErrWalletNotFound         = errors.New("wallet not found")
	ErrClientNotFound         = errors.New("client not found")
	ErrWalletOperationInvalid = errors.New("wallet operation invalid")
)
