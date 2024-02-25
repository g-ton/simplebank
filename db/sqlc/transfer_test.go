package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	from_account := CreateRandomAccount(t)
	to_account := CreateRandomAccount(t)

	amountInt, _ := faker.RandomInt(0, 10000, 1)
	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		// minus to have possibly negative values
		Amount: int64(amountInt[0]) - 20,
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	trans1 := createRandomTransfer(t)
	trans2, err := testQueries.GetTransfer(context.Background(), trans1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.Equal(t, trans1.ID, trans2.ID)
	require.Equal(t, trans1.FromAccountID, trans2.FromAccountID)
	require.Equal(t, trans1.ToAccountID, trans2.ToAccountID)
	require.Equal(t, trans1.Amount, trans2.Amount)
	require.WithinDuration(t, trans1.CreatedAt, trans2.CreatedAt, time.Second)
}

func TestUpdateTransfer(t *testing.T) {
	trans1 := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     trans1.ID,
		Amount: trans1.Amount + 1,
	}
	trans2, err := testQueries.UpdateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.Equal(t, arg.ID, trans2.ID)
	require.Equal(t, arg.Amount, trans2.Amount)
	require.Equal(t, trans1.FromAccountID, trans2.FromAccountID)
	require.Equal(t, trans1.ToAccountID, trans2.ToAccountID)
	require.WithinDuration(t, trans1.CreatedAt, trans2.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	trans1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), trans1.ID)
	require.NoError(t, err)

	trans2, err := testQueries.GetTransfer(context.Background(), trans1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, trans2)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}
