// Package users
// Author: Duc Hung Ho @kyeranyo
// Description: users repository implementation
package users

import (
	"context"
	"errors"

	"swclabs/swipecore/internal/core/domain"
	"swclabs/swipecore/pkg/db"
)

type Users struct {
	db db.IDatabase
}

func New(conn db.IDatabase) IUserRepository {
	return useCache(&Users{conn})
}

var _ IUserRepository = (*Users)(nil)

// GetByEmail implements IUserRepository.
func (usr *Users) GetByEmail(ctx context.Context, email string) (*domain.Users, error) {
	rows, err := usr.db.Query(ctx, selectByEmail, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Insert implements IUserRepository.
func (usr *Users) Insert(ctx context.Context, _usr domain.Users) error {
	return usr.db.SafeWrite(
		ctx,
		insertIntoUsers,
		_usr.Email, _usr.PhoneNumber, _usr.FirstName, _usr.LastName, _usr.Image,
	)
}

// Info implements IUserRepository.
func (usr *Users) Info(ctx context.Context, email string) (*domain.UserSchema, error) {
	rows, err := usr.db.Query(ctx, selectUserInfo, email)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.UserSchema](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SaveInfo implements IUserRepository.
func (usr *Users) SaveInfo(ctx context.Context, user domain.Users) error {
	if user.Email == "" {
		return errors.New("missing key: email ")
	}
	if user.FirstName != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.LastName != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	}
	if user.Image != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	if user.PhoneNumber != "" {
		if err := usr.db.SafeWrite(
			ctx, updateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	}
	return nil
}

// UpdateProperties implements IUserRepository.
func (usr *Users) UpdateProperties(
	ctx context.Context, query string, user domain.Users) error {
	switch query {
	case updateUsersLastname:
		if err := usr.db.SafeWrite(ctx,
			updateUsersLastname, user.LastName, user.Email,
		); err != nil {
			return err
		}
	case updateUsersFirstname:
		if err := usr.db.SafeWrite(ctx,
			updateUsersFirstname, user.FirstName, user.Email,
		); err != nil {
			return err
		}
	case updateUsersPhoneNumber:
		if err := usr.db.SafeWrite(ctx,
			updateUsersPhoneNumber, user.PhoneNumber, user.Email,
		); err != nil {
			return err
		}
	case updateUsersImage:
		if err := usr.db.SafeWrite(ctx,
			updateUsersImage, user.Image, user.Email,
		); err != nil {
			return err
		}
	}
	return errors.New("unknown :" + query)
}

// OAuth2SaveInfo implements IUserRepository.
func (usr *Users) OAuth2SaveInfo(ctx context.Context, user domain.Users) error {
	return usr.db.SafeWrite(
		ctx, insertUsersConflict, user.Email, user.PhoneNumber,
		user.FirstName, user.LastName, user.Image,
	)
}

// GetByPhone implements IUserRepository.
func (usr *Users) GetByPhone(ctx context.Context, nPhone string) (*domain.Users, error) {
	rows, err := usr.db.Query(ctx, selectByPhone, nPhone)
	if err != nil {
		return nil, err
	}
	user, err := db.CollectOneRow[domain.Users](rows)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
