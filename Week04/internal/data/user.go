package data

import (
	"github.com/pkg/errors"
)

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Nickname string `db:"nickname"`
}

type UserDao struct {
	conn *DBConn
}

func NewUserDao(conn *DBConn) *UserDao {
	return &UserDao{conn: conn}
}

func (dao UserDao) QueryById(uid int) (*User, error) {
	user := &User{}
	err := dao.conn.QueryRow("select id, username, password, nickname from `fake_user` where id=?", uid).Scan(&user.Id,
		&user.Username, &user.Password, &user.Nickname)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}
