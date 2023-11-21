package repo

import (
	"example/swiftcart/internal/model"
	"example/swiftcart/internal/schema"
	"example/swiftcart/pkg/db"
	"example/swiftcart/pkg/db/queries"

	"gorm.io/gorm"
)

type Users struct {
	conn *gorm.DB
	data *model.User
}

func NewUsers() IUsers {
	_conn, err := db.Connection()
	if err != nil {
		panic(err)
	}
	return &Users{
		conn: _conn,
		data: &model.User{},
	}
}

func (usr *Users) GetByEmail(email string) (*model.User, error) {
	if err := usr.conn.Table("users").Where("email = ?", email).First(usr.data).Error; err != nil {
		return nil, err
	}
	return usr.data, nil
}

func (usr *Users) Insert(_usr *model.User) error {
	return usr.conn.Exec(
		queries.INSERT_INTO_USERS,
		_usr.Email,
		_usr.PhoneNumber,
		_usr.FirstName,
		_usr.LastName,
		_usr.Image).Error
}

func (usr *Users) Info(email string) (*schema.UserInfo, error) {
	data := new(schema.UserInfo)
	if err := usr.conn.Raw(queries.SELECT_USER_INFO, email).Scan(data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (usr *Users) SaveInfo(user *model.User) error {
	return nil
}