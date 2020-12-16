package dto

type UserDTO struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
}
