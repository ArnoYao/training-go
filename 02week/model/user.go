package model

type User struct {
	Id   int64
	Name string
}

type IUserDB interface {
	Get(id int64) (*User, error)
}
