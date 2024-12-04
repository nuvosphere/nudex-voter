package withdrawal

import "github.com/nuvosphere/nudex-voter/internal/db"

func CheckTx(task *db.WithdrawalTask) (bool, error) {
	return true, nil
}

func CheckBalance(task *db.WithdrawalTask) (bool, error) {
	return true, nil
}
