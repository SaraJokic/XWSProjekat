package domain

type AuthStore interface {
	DeleteAll()
	Insert(user *User) error
	GetAll() ([]*User, error)
	ValidateUsernameAndPassword(username string, password string) (*User, error)
	Delete(id string) error
	GetByUsername(username string) (*User, error)
}
