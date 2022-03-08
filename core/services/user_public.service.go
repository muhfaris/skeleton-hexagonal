package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/muhfaris/skeleton-hexagonal/core/entities/account/model"
	svcmodel "github.com/muhfaris/skeleton-hexagonal/core/entities/account/service"
	"github.com/muhfaris/skeleton-hexagonal/core/repository"
	"github.com/muhfaris/skeleton-hexagonal/transport/structures"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserPublicService interface {
	Login(ctx context.Context, params *structures.LoginRead) (svcmodel.AccountResponse, error)
	SignUp(ctx context.Context, params *structures.SignUpRead) error
}

type userPublicService struct {
	repo        repository.UserPublicRepository
	accountRepo repository.AccountRepository
}

func NewUserPublicnService(repo repository.UserPublicRepository, accountRepo repository.AccountRepository) UserPublicService {
	return &userPublicService{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (service *userPublicService) Login(ctx context.Context, params *structures.LoginRead) (svcmodel.AccountResponse, error) {
	result := <-service.accountRepo.FindByEmail(ctx, params.Username)
	if result.Error != nil {
		return svcmodel.AccountResponse{}, result.Error
	}

	account, ok := result.Data.(model.Account)
	if !ok {
		return svcmodel.AccountResponse{}, fmt.Errorf("error: parse data account to object")
	}

	if err := account.ComparePassword(params.Password); err != nil {
		return svcmodel.AccountResponse{}, err
	}

	return *account.Response(), nil
}

func (service *userPublicService) SignUp(ctx context.Context, params *structures.SignUpRead) error {
	result := <-service.accountRepo.FindByEmail(ctx, params.Email)
	if result.Error != nil && result.Error != pgx.ErrNoRows && result.Error != mongo.ErrNoDocuments {
		return result.Error
	}

	if result.Error == pgx.ErrNoRows || result.Error == mongo.ErrNoDocuments {
		account := model.CreateAccount(params)
		if err := account.GenerateHashPassword(); err != nil {
			return err
		}

		if result := <-service.repo.SignUp(ctx, account); result.Error != nil {
			return result.Error
		}

		return nil
	}

	_, ok := result.Data.(model.Account)
	if !ok {
		return fmt.Errorf("error: parse data account to object")
	}

	return fmt.Errorf("email already exists")
}
