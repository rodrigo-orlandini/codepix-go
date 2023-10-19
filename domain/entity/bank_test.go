package entity_test

import (
	"testing"

	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestEntity_NewBank(t *testing.T) {
	code := "001"
	name := "Banco Santander"
	bank, validationError := entity.NewBank(code, name)

	require.Nil(t, validationError)
	require.NotEmpty(t, uuid.FromStringOrNil(bank.ID))
	require.Equal(t, bank.Code, code)
	require.Equal(t, bank.Name, name)

	_, validationError = entity.NewBank("", "")
	require.NotNil(t, validationError)
}
