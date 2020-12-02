package service

import (
	"github.com/pkg/errors"
	"github.com/x-debug/Go-000/Week02/dao"
)

func GetUser(uid uint64) (*dao.User, error) {
	user, err := dao.FindById(uid)

	if err != nil {
		return nil, errors.WithMessage(err, "get user error")
	}

	return user, nil
}
