//+build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/x-debug/Go-000/Week04/internal/biz"
	"github.com/x-debug/Go-000/Week04/internal/data"
)

func CreateUserCase() *biz.UserCase {
	wire.Build(biz.NewUserCase, data.NewUserDao, data.NewDB)
	return nil
}
