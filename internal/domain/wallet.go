package domain

type Operation string

const (
	Deposit  Operation = "DEPOSIT"
	Withdraw Operation = "WITHDRAW"
)

var validOperation = map[Operation]struct{}{
	Deposit:  {},
	Withdraw: {},
}

func NewOperation(o string) (Operation, error) {
	newOperation := Operation(o)

	if _, ok := validOperation[newOperation]; !ok {
		return "", ErrWalletOperationInvalid
	}

	return newOperation, nil
}

type Wallet struct {
	ID      string
	Balance int64
}
