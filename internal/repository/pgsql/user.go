package pgsql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"github.com/rahmat412/go-microservice-template/database/service_template/public/model"
	"github.com/rahmat412/go-microservice-template/database/service_template/public/table"
	"github.com/rahmat412/go-microservice-template/internal/helper/customerror"
	toolboxError "github.com/rahmat412/go-toolbox/error"
)

type UserRepositoryProvider interface {
	GetUserByID(ctx context.Context, id int) (*model.Users, error)
	CreateUser(ctx context.Context, user *model.Users) (*model.Users, error)
	UpdateUser(ctx context.Context, user *model.Users) (*model.Users, error)
	DeleteUserByID(ctx context.Context, id int) error
}

type userRepositoryImplementation struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryProvider {
	return &userRepositoryImplementation{db: db}
}

func (u *userRepositoryImplementation) GetUserByID(ctx context.Context, id int) (*model.Users, error) {
	statement := table.Users.SELECT(
		table.Users.ID,
		table.Users.Username,
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.BirthDate,
		table.Users.Email,
		table.Users.Password,
		table.Users.IsActive,
		table.Users.CreatedAt,
	).WHERE(table.Users.ID.EQ(postgres.Int(int64(id))))

	var user model.Users
	err := statement.QueryContext(ctx, u.db, &user)
	if err != nil && !errors.Is(err, qrm.ErrNoRows) {
		return nil, customerror.ErrorInternalServer.WithCause(err).WithLocator(toolboxError.WhereAmI()).WithStackTrace()
	} else if errors.Is(err, qrm.ErrNoRows) {
		return nil, nil
	}

	return &user, nil
}

func (u *userRepositoryImplementation) CreateUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	statement := table.Users.INSERT(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.BirthDate,
		table.Users.CreatedAt,
		table.Users.IsActive,
	).MODEL(user)

	_, err := statement.ExecContext(ctx, u.db)
	if err != nil {
		return nil, customerror.ErrorInternalServer.WithCause(err).WithLocator(toolboxError.WhereAmI()).WithStackTrace()
	}

	return user, nil
}

func (u *userRepositoryImplementation) UpdateUser(ctx context.Context, user *model.Users) (*model.Users, error) {
	statement := table.Users.UPDATE(
		table.Users.FirstName,
		table.Users.LastName,
		table.Users.BirthDate,
		table.Users.IsActive,
	).SET(
		table.Users.FirstName.SET(postgres.String(user.FirstName)),
		table.Users.LastName.SET(postgres.String(*user.LastName)),
		table.Users.BirthDate.SET(postgres.Date(user.BirthDate.Date())),
		table.Users.IsActive.SET(postgres.Bool(*user.IsActive)),
	).WHERE(table.Users.ID.EQ(postgres.Int(int64(user.ID))))

	_, err := statement.ExecContext(ctx, u.db)
	if err != nil {
		return nil, customerror.ErrorInternalServer.WithCause(err).WithLocator(toolboxError.WhereAmI()).WithStackTrace()
	}

	return user, nil
}

func (u *userRepositoryImplementation) DeleteUserByID(ctx context.Context, id int) error {
	statement := table.Users.DELETE().WHERE(table.Users.ID.EQ(postgres.Int(int64(id))))

	_, err := statement.ExecContext(ctx, u.db)
	if err != nil {
		return customerror.ErrorInternalServer.WithCause(err).WithLocator(toolboxError.WhereAmI()).WithStackTrace()
	}

	return nil
}
