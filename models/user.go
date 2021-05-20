package models

type User struct {
	ID       string `json:"id" gorm:"unique;not null"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password"`
}

type UserRepo interface {
	Save(user *User) (*User, error)
	FindAll() ([]User, error)
	Delete(user *User) error
	Migrate() error
}
