package usecase

import "context"

func (u *WalletUsecase) GetBalance(ctx context.Context, id string) (int64, error) {
	wallet, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return 0, err
	}
	return wallet.Balance, nil
}
