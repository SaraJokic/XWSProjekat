package domain

type AuthStore interface {
	DeleteAll()
	Insert(user *User) error
	GetAll() ([]*User, error)
}
