package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func createRandomEntries(t *testing.T, account Account) Entries {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entries, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)

	require.Equal(t, arg.AccountID, entries.AccountID)
	require.Equal(t, arg.Amount, entries.Amount)

	require.NotZero(t, entries.ID)
	require.NotZero(t, entries.Amount)

	return entries
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}

func TestGetEntries(t *testing.T) {
	account := createRandomAccount(t)
	entries1 := createRandomEntries(t, account)
	entries2, err := testQueries.GetEntries(context.Background(), entries1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entries2)

	require.Equal(t, entries1.ID, entries2.ID)
	require.Equal(t, entries1.AccountID, entries2.AccountID)
	require.Equal(t, entries1.Amount, entries2.Amount)

	require.WithinDuration(t, entries1.CreateAt, entries2.CreateAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntries(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
