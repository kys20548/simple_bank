package db

import (
	"context"
	"github.com/kys20548/simple_bank/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: createRandomAccount(t).ID,
		Amount:    util.RandomMoney(),
	}
	entry1, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, arg.AccountID, entry1.AccountID)
	require.Equal(t, arg.Amount, entry1.Amount)

	return entry1
}
func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}
