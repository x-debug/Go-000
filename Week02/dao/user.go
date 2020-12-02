package dao

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/x-debug/Go-000/Week02/dao/db"
)

type User struct {
	Username string

	Password string
}

func (u *User) String() string  {
	return fmt.Sprintf("Username: %s", u.Username)
}

func FindById(uid uint64) (*User, error) {
	row, err := db.GetByPrimary(uid)

	if err != nil {
		return nil, errors.Wrap(err, "find user error")
	}

	return &User{Username: row.Col1, Password: row.Col2}, nil
}
