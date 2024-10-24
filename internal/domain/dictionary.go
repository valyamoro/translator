package domain

type Dictionary struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"max=255"`
	UserID      int64  `json:"user_id" validate:"user_exists"`
}
