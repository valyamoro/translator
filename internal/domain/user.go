package domain

type User struct {
	ID           int64        `json:"id"`
	Username     string       `json:"username" validate:"required,min=3,max=50,exists_user"`
	Password     string       `json:"password" validate:"required,min=8,passwd"`
	Dictionaries []Dictionary `json:"dictionaries"`
}
