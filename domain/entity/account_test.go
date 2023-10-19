package entity_test

import (
	"testing"

	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestEntity_NewAccount(t *testing.T) {
	code := "001"
	name := "Banco Itaú"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "John Doe"
	account, validationError := entity.NewAccount(bank, ownerName, accountNumber)

	require.Nil(t, validationError)
	require.NotEmpty(t, uuid.FromStringOrNil(account.ID))
	require.Equal(t, account.Number, accountNumber)
	require.Equal(t, account.Bank.ID, bank.ID)

	_, validationError = entity.NewAccount(bank, "", ownerName)
	require.NotNil(t, validationError)
}
