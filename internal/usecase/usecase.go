package usecase

import (
	"context"

	"wallet/internal/domain"
)

//go:generate mockgen -source=usecase.go -destination=../../mocks/usecase/wallet_repo_mock.go -package=usecase
type WalletRepository interface {
	GetByID(ctx context.Context, id string) (*domain.Wallet, error)
	UpdateBalance(ctx context.Context, id string, amount int64) error
}

type WalletUsecase struct {
	repo WalletRepository
}

func NewWalletUsecase(repo WalletRepository) *WalletUsecase {
	return &WalletUsecase{repo: repo}
}
