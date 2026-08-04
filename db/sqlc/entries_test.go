package db

import (
	"context"
	"simplebank/utils"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: 1,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := createRandomEntry(t)
	entryID := entry.ID
	entry, err := testQueries.GetEntry(context.Background(), entryID)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entryID, entry.ID)

	require.NotZero(t, entry.ID)
}

func TestListEntries(t *testing.T) {
	_ = createRandomEntry(t)
	_ = createRandomEntry(t)

	arg := ListEntriesParams{
		AccountID: 1,
		Limit:     2,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entries)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

	require.Len(t, entries, int(arg.Limit))
}
