package biz

import (
	"github.com/x-debug/Go-000/Week04/internal/data"
	"github.com/x-debug/Go-000/Week04/internal/dto"
)

type UserCase struct {
	userDao *data.UserDao
}

func NewUserCase(dao *data.UserDao) *UserCase {
	return &UserCase{userDao: dao}
}

func (uc *UserCase) FindById(uid int) (*dto.UserDTO, error) {
	user, err := uc.userDao.QueryById(uid)

	if err != nil {
		return nil, err
	}

	return &dto.UserDTO{Username: user.Username, Id: user.Id}, nil
}
