package entity_test

import (
	"testing"

	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestEntity_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, _ := entity.NewBank(code, name)

	accountNumber := "12345-6"
	ownerName := "John Doe"
	account, _ := entity.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "johndoe@example.com"
	pixKey, _ := entity.NewPixKey(kind, key, account)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, validationError := entity.NewPixKey(kind, key, account)
	require.Nil(t, validationError)

	_, validationError = entity.NewPixKey("nome", key, account)
	require.NotNil(t, validationError)
}
