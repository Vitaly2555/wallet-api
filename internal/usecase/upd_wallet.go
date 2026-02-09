package usecase

import (
	"context"
	"wallet/internal/domain"
)

func (u *WalletUsecase) Operate(ctx context.Context, id string, op domain.Operation, amount int64) error {
	if op == domain.Withdraw {
		amount = -amount
	}
	return u.repo.UpdateBalance(ctx, id, amount)
}
