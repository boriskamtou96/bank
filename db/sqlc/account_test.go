package db

import (
	"context"
	"simplebank/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	accountID := int64(1)
	account, err := testQueries.GetAccount(context.Background(), accountID)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, accountID, account.ID)

	require.NotZero(t, account.ID)
}

func TestDeleteAccount(t *testing.T) {
	accountID := int64(2)

	err := testQueries.DeleteAccount(context.Background(), accountID)

	require.NoError(t, err)
}

func TestGetListAccounts(t *testing.T) {
	arg := ListAccountsParams{
		Limit:  10,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
