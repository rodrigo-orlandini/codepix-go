package entity_test

import (
	"testing"

	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	code := "001"
	name := "Banco Inter"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "John Doe"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "56789-0"
	ownerName = "Joana Doe"
	accountDestination, _ := entity.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "joanadoe@example.com"
	pixKey, _ := entity.NewPixKey(kind, key, accountDestination)

	require.NotEqual(t, account.ID, accountDestination.ID)

	amount := 100.0
	statusTransaction := "pending"
	transaction, validationError := entity.NewTransaction(account, amount, pixKey, "Some interesting description")

	require.Nil(t, validationError)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "Some interesting description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount, _ := entity.NewPixKey(kind, key, account)

	_, validationError = entity.NewTransaction(account, amount, pixKeySameAccount, "Some interesting description")
	require.NotNil(t, validationError)

	_, validationError = entity.NewTransaction(account, 0, pixKey, "Some interesting description")
	require.NotNil(t, validationError)
}

func TestEntity_ChangeStatusOfATransaction(t *testing.T) {
	code := "001"
	name := "Banco Inter"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "John Doe"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	accountNumberDestination := "56789-0"
	ownerName = "Joana Doe"
	accountDestination, _ := entity.NewAccount(bank, accountNumberDestination, ownerName)

	kind := "email"
	key := "joanadoe@example.com"
	pixKey, _ := entity.NewPixKey(kind, key, accountDestination)

	amount := 100.0
	transaction, _ := entity.NewTransaction(account, amount, pixKey, "Some interesting description")

	transaction.Complete()
	require.Equal(t, transaction.Status, entity.TransactionCompleted)

	transaction.Cancel("Error")
	require.Equal(t, transaction.Status, entity.TransactionError)
	require.Equal(t, transaction.CancelDescription, "Error")
}
