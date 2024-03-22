package db

import (
	"context"
	"testing"

	"github.com/dev-vickie/Bank-Model-DevOps/db/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomAmount(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)
}
