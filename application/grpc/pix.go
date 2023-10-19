package grpc

import (
	"context"

	"github.com/rodrigo-orlandini/codepix-go/application/grpc/pb"
	"github.com/rodrigo-orlandini/codepix-go/application/usecase"
)

type PixGRPCService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func NewPixGRPCService(usecase usecase.PixUseCase) *PixGRPCService {
	return &PixGRPCService{
		PixUseCase: usecase,
	}
}

func (service *PixGRPCService) RegisterPixKey(context context.Context, data *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	pixKey, err := service.PixUseCase.RegisterKey(data.Kind, data.Key, data.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "Pix Key not created.",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     pixKey.ID,
		Status: "Created",
	}, nil
}

func (service *PixGRPCService) Find(context context.Context, data *pb.PixKey) (*pb.PixKeyInfo, error) {
	pixKey, err := service.PixUseCase.FindKey(data.Key, data.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   pixKey.ID,
		Kind: pixKey.Kind,
		Key:  pixKey.Key,
		Account: &pb.Account{
			AccountId:     pixKey.AccountID,
			AccountNumber: pixKey.Account.Number,
			BankId:        pixKey.Account.BankID,
			BankName:      pixKey.Account.Bank.Name,
			OwnerName:     pixKey.Account.OwnerName,
			CreatedAt:     pixKey.Account.CreatedAt.String(),
		},
		CreatedAt: pixKey.CreatedAt.String(),
	}, nil
}
