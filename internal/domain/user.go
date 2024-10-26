package domain

type User struct {
	ID           int64        `json:"id"`
	Username     string       `json:"username"`
	Password     string       `json:"password"`
	Dictionaries []Dictionary `json:"dictionaries"`
}

func (u User) Equals(other User) bool {
	if u.ID != other.ID || u.Username != other.Username {
		return false 
	}

	if len(u.Dictionaries) != len(other.Dictionaries) {
		return false 
	}

	for i := range u.Dictionaries {
		if u.Dictionaries[i] != other.Dictionaries[i] {
			return false 
		}
	}

	return true
}
