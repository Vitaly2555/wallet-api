package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"wallet/internal/domain"
	"wallet/mocks/usecase"
)

type usecaseMocks struct {
	walletRepo *usecase.MockWalletRepository
}

func setup(t *testing.T) (*WalletUsecase, *usecaseMocks) {
	ctrl := gomock.NewController(t)
	t.Cleanup(ctrl.Finish)

	mocks := &usecaseMocks{
		walletRepo: usecase.NewMockWalletRepository(ctrl),
	}

	return &WalletUsecase{
		repo: mocks.walletRepo,
	}, mocks
}

func Test_Operate(t *testing.T) {
	ctx := context.TODO()

	var (
		testID = "test-id"

		fakeErr = errors.New("fake err")
	)

	tests := []struct {
		name string

		id        string
		operation domain.Operation
		amount    int64

		setMocks func(mock *usecaseMocks)

		expErr string
	}{
		{
			name:      "success",
			id:        testID,
			operation: domain.Deposit,
			amount:    1000,
			setMocks: func(mock *usecaseMocks) {
				mock.walletRepo.
					EXPECT().
					UpdateBalance(ctx, testID, int64(1000)).
					Return(nil)
			},
			expErr: "",
		},
		{
			name:      "success_with_withdraw",
			id:        testID,
			operation: domain.Withdraw,
			amount:    1000,
			setMocks: func(mock *usecaseMocks) {
				mock.walletRepo.
					EXPECT().
					UpdateBalance(ctx, testID, int64(-1000)).
					Return(nil)
			},
			expErr: "",
		},
		{
			name:      "err",
			id:        testID,
			operation: domain.Deposit,
			amount:    1000,
			setMocks: func(mock *usecaseMocks) {
				mock.walletRepo.
					EXPECT().
					UpdateBalance(ctx, testID, int64(1000)).
					Return(fakeErr)
			},
			expErr: "fake err",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, mocks := setup(t)

			if tt.setMocks != nil {
				tt.setMocks(mocks)
			}

			err := s.Operate(ctx, tt.id, tt.operation, tt.amount)
			if err != nil {
				require.Equal(t, tt.expErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_GetBalance(t *testing.T) {
	ctx := context.TODO()

	var (
		testID = "test-id"

		fakeErr = errors.New("fake err")
	)

	tests := []struct {
		name string

		id string

		setMocks func(mock *usecaseMocks)

		expBalance int64
		expErr     string
	}{
		{
			name: "success",
			id:   testID,
			setMocks: func(mock *usecaseMocks) {
				mock.walletRepo.
					EXPECT().
					GetByID(ctx, testID).
					Return(&domain.Wallet{
						ID:      testID,
						Balance: 1000,
					}, nil)
			},
			expBalance: int64(1000),
			expErr:     "",
		},
		{
			name: "err",
			id:   testID,
			setMocks: func(mock *usecaseMocks) {
				mock.walletRepo.
					EXPECT().
					GetByID(ctx, testID).
					Return(nil, fakeErr)
			},
			expBalance: 0,
			expErr:     "fake err",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s, mocks := setup(t)

			if tt.setMocks != nil {
				tt.setMocks(mocks)
			}

			balance, err := s.GetBalance(ctx, tt.id)
			if err != nil {
				require.Equal(t, tt.expErr, err.Error())
			} else {
				require.Equal(t, tt.expBalance, balance)
				require.NoError(t, err)
			}
		})
	}
}
