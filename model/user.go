package model

// User - Data Model To Store user Data in
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u User) Valid() bool {
	return u.ID > 0 && u.FirstName != "" && u.LastName != "" && u.Email != ""
}
