//+build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/x-debug/Go-000/Week04/internal/biz"
	"github.com/x-debug/Go-000/Week04/internal/data"
)

func InitializeDB() (*data.DBConn, func(), error) {
	panic(wire.Build(data.NewDB, data.NewConf))
}

func createUserCase() (*biz.UserCase, func(), error) {
	panic(wire.Build(biz.NewUserCase, data.NewUserDao, data.NewDB, data.NewConf))
}

func GetUserCase() *biz.UserCase {
	userCase, _, _ := createUserCase()
	return userCase
}
