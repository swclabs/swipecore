// Package users users repository implementation
package users

import (
	"context"

	"swclabs/swix/internal/core/domain/entity"
	"swclabs/swix/internal/core/domain/model"
	"swclabs/swix/pkg/infra/cache"
	"swclabs/swix/pkg/infra/db"
)

// Users repository implementation
type Users struct {
	db db.IDatabase
}

// New creates a new instance of IUserRepository.
func New(conn db.IDatabase) IUserRepository {
	return &Users{conn}
}

// Init initializes the Users object with database and redis connection
func Init(conn db.IDatabase, cache cache.ICache) IUserRepository {
	return useCache(cache, New(conn))
}

var _ IUserRepository = (*Users)(nil)

// GetByID implements IUserRepository.
func (usr *Users) GetByID(ctx context.Context, id int64) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByID, id)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByEmail implements IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr entity.Users) error {
	return usr.db.SafeWrite(
		ctx,
		insertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*model.Users, error) {
	rows, err := usr.db.Query(ctx, selectUserInfo, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[model.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Save implements IUserRepository.
func (usr *Users) Save(ctx context.Context, user entity.Users) error {
	return usr.db.SafeWrite(ctx, updateInfo,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Image,
		user.PhoneNumber,
	)
}

// OAuth2SaveInfo implements IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user entity.Users) error {
	return usr.db.SafeWrite(
		ctx, insertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// GetByPhone implements IUserRepository.
func (usr *Users) GetByPhone(ctx context.Context, nPhone string) (*entity.Users, error) {
	rows, err := usr.db.Query(ctx, selectByPhone, nPhone)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[entity.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
