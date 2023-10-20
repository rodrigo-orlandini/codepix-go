package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/rodrigo-orlandini/codepix-go/application/usecase"
	"github.com/rodrigo-orlandini/codepix-go/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) *usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDatabase{Database: database}
	transactionRepository := repository.TransactionRepositoryDatabase{Database: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: transactionRepository,
		PixKeyRepository:      pixRepository,
	}

	return &transactionUseCase
}
