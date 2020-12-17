//+build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/x-debug/Go-000/Week04/internal/biz"
	"github.com/x-debug/Go-000/Week04/internal/data"
)

var dbProviders = wire.NewSet(data.NewDB, data.NewConf)

func InitializeDB() (*data.DBConn, func(), error) {
	panic(wire.Build(data.NewDB, data.NewConf))
}

func createUserCase() (*biz.UserCase, func(), error) {
	panic(wire.Build(biz.NewUserCase, wire.NewSet(dbProviders, data.NewUserDao)))
}

func GetUserCase() *biz.UserCase {
	userCase, _, _ := createUserCase()
	return userCase
}
