package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/dev-vickie/Bank-Model-DevOps/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
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

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	//create account - each test should create account so that it can be independent of the rest
	newAccount := createRandomAccount(t)
	returnedAcc, err := testQueries.GetAccount(context.Background(), newAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, returnedAcc)

	require.Equal(t, newAccount.ID, returnedAcc.ID)
	require.Equal(t, newAccount.Owner, returnedAcc.Owner)
	require.Equal(t, newAccount.Currency, returnedAcc.Currency)
	require.Equal(t, newAccount.Balance, returnedAcc.Balance)
	require.WithinDuration(t, newAccount.CreatedAt, returnedAcc.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T){
	newAccount := createRandomAccount(t)

	arg:= UpdateAccountParams{
		ID: newAccount.ID,
		Balance: util.RandomAmount(),
	}

	returnedAcc,err := testQueries.UpdateAccount(context.Background(),arg)
	require.NoError(t,err)
	require.NotEmpty(t,returnedAcc)

	require.Equal(t, newAccount.ID, returnedAcc.ID)
	require.Equal(t, newAccount.Owner, returnedAcc.Owner)
	require.Equal(t, newAccount.Currency, returnedAcc.Currency)
	require.Equal(t, arg.Balance, returnedAcc.Balance)
	require.WithinDuration(t, newAccount.CreatedAt, returnedAcc.CreatedAt, time.Second)

}

func TestDeletAccount(t *testing.T){
	newAccount := createRandomAccount(t)
	
	err := testQueries.DeleteAccount(context.Background(),newAccount.ID)
	require.NoError(t,err)

	account2,err := testQueries.GetAccount(context.Background(),newAccount.ID)
	require.Error(t,err)
	require.Empty(t,account2)
	require.EqualError(t,err,sql.ErrNoRows.Error())
}

func TestListAllAccounts(t *testing.T){
	for i:=0;i < 10; i++{
		createRandomAccount(t)
	}

	args:= ListAccountsParams{
		Limit: 5,
		Offset: 5,
	}

	accounts,err := testQueries.ListAccounts(context.Background(),args)
	require.NoError(t,err)
	require.Len(t,accounts,5)

	for _, account :=range accounts{
		require.NotEmpty(t,account)
	}
}