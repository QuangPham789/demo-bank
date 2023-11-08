package db

import (
	"context"
	"simple-bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, _ := testQueries.CreateAccount(context.Background(), arg)

	return account
}
func TestQueries_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, arg.Owner)
	require.Equal(t, arg.Balance, arg.Balance)
	require.Equal(t, arg.Currency, arg.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

}
