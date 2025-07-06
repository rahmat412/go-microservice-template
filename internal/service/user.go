package service

import (
	"context"
	"time"

	"github.com/rahmat412/go-microservice-template/database/service_template/public/model"
	"github.com/rahmat412/go-microservice-template/internal/dto"
	"github.com/rahmat412/go-microservice-template/internal/helper/customerror"
	"github.com/rahmat412/go-microservice-template/internal/helper/date"
	"github.com/rahmat412/go-microservice-template/internal/repository/pgsql"
	"github.com/rahmat412/go-toolbox/logging"
)

type UserServiceProvider interface {
	CreateUser(ctx context.Context, user dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	GetUserByID(ctx context.Context, id int) (*dto.GetUserByIDResponse, error)
	UpdateUser(ctx context.Context, id int, user *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, id int) error
}

type userServiceImplementation struct {
	repo    pgsql.UserRepositoryProvider
	log     *logging.Logger
	timeNow func() time.Time
}

func (u userServiceImplementation) CreateUser(ctx context.Context, user dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	now := u.timeNow()
	parseTime, err := date.ParseStringToDate(ctx, user.BirthDate)
	if err != nil {
		return nil, err
	}

	newUser := &model.Users{
		FirstName: user.FirstName,
		LastName:  &user.LastName,
		BirthDate: &parseTime,
		IsActive:  &user.IsActive,
		CreatedAt: &now,
	}

	createdUser, err := u.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponse{
		ID:        int(createdUser.ID),
		FirstName: createdUser.FirstName,
		LastName:  *createdUser.LastName,
		BirthDate: createdUser.BirthDate.Format(time.RFC3339),
		CreatedAt: createdUser.CreatedAt.Format(time.RFC3339),
		IsActive:  *createdUser.IsActive,
	}, nil
}

func (u userServiceImplementation) GetUserByID(ctx context.Context, id int) (*dto.GetUserByIDResponse, error) {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &dto.GetUserByIDResponse{
		ID:        int(user.ID),
		FirstName: user.FirstName,
		LastName:  *user.LastName,
		BirthDate: user.BirthDate.Format(time.RFC3339),
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		IsActive:  *user.IsActive,
	}, nil
}

func (u userServiceImplementation) UpdateUser(ctx context.Context, id int, user *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	currentUser, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if currentUser == nil {
		return nil, customerror.ErrorUserNotFound
	}

	parseDate, err := date.ParseStringToDate(ctx, user.BirthDate)
	if err != nil {
		return nil, err
	}

	updatedUser := &model.Users{
		ID:        int32(id),
		FirstName: user.FirstName,
		LastName:  &user.LastName,
		BirthDate: &parseDate,
		IsActive:  nil,
		CreatedAt: nil,
	}

	updatedUser, err = u.repo.UpdateUser(ctx, updatedUser)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserResponse{
		ID:        int(updatedUser.ID),
		FirstName: updatedUser.FirstName,
		LastName:  *updatedUser.LastName,
		BirthDate: user.BirthDate,
	}, nil

}

func (u userServiceImplementation) DeleteUser(ctx context.Context, id int) error {
	user, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return err
	}

	if user == nil {
		return customerror.ErrorUserNotFound
	}

	err = u.repo.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(repo pgsql.UserRepositoryProvider, log *logging.Logger) UserServiceProvider {
	return &userServiceImplementation{
		repo: repo,
		log:  log,
		timeNow: func() time.Time {
			return time.Now()
		},
	}
}
