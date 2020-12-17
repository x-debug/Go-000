//+build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/x-debug/Go-000/Week04/internal/biz"
	"github.com/x-debug/Go-000/Week04/internal/data"
)

func InitializeDB() (*data.DBConn, func()) {
	wire.Build(data.NewDB, data.NewConf)
	return nil, nil
}

func createUserCase() (*biz.UserCase, func()) {
	wire.Build(biz.NewUserCase, data.NewUserDao, data.NewDB, data.NewConf)
	return nil, nil
}

func GetUserCase() *biz.UserCase {
	userCase, _ := createUserCase()
	return userCase
}
