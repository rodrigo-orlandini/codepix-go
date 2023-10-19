package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/domain/entity"
)

// type ITransactionRepository interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	Find(id string) (*Transaction, error)
// }

type TransactionRepositoryDatabase struct {
	Database *gorm.DB
}

func (repository TransactionRepositoryDatabase) Register(transaction *entity.Transaction) error {
	databaseError := repository.Database.Create(transaction).Error

	if databaseError != nil {
		return databaseError
	}

	return nil
}

func (repository TransactionRepositoryDatabase) Save(transaction *entity.Transaction) error {
	databaseError := repository.Database.Save(transaction).Error

	if databaseError != nil {
		return databaseError
	}

	return nil
}

func (repository TransactionRepositoryDatabase) Find(id string) (*entity.Transaction, error) {
	var transaction entity.Transaction

	repository.Database.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, errors.New("Transaction was not found.")
	}

	return &transaction, nil
}
