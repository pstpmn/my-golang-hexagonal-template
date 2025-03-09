package domain

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

func (u *User) SetID(id string) {
	u.ID = id
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetIsActive(isActive bool) {
	u.IsActive = isActive
}

func NewUser(id, name, email string, isActive bool) (*User, error) {
	u := &User{}
	u.SetID(id)
	u.SetName(name)
	u.SetEmail(email)
	u.SetIsActive(isActive)
	return u, nil
}
