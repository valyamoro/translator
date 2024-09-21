package domain

type Dictionary struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int64  `json:"user_id"`
}

type UpdateDictionaryInput struct {
	ID          *int64  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	UserID      *int64  `json:"user_id"`
}
