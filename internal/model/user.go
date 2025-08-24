package model

type UserID uint

func NewUserID(id uint) (UserID, error) {
	return UserID(id), nil
}

type Email string

func NewEmail(email string) (Email, error) {
	return Email(email), nil
}

type Username string

func NewUsername(username string) (Username, error) {
	return Username(username), nil
}

type PasswordHash string

func NewPasswordHash(passwordHash string) (PasswordHash, error) {
	return PasswordHash(passwordHash), nil
}

type User struct {
	ID       UserID
	Email    Email
	Username Username
	Role     Role
}

func (u User) GetEmail() string {
	return string(u.Email)
}

func NewUser(id uint, email, username, role string, permissions []string) (*User, error) {
	uid, err := NewUserID(uint(id))
	if err != nil {
		return &User{}, err
	}
	uemail, err := NewEmail(email)
	if err != nil {
		return &User{}, err
	}
	uusername, err := NewUsername(username)
	if err != nil {
		return &User{}, err
	}
	urole, err := NewRole(role, permissions)
	if err != nil {
		return &User{}, err
	}

	return &User{ID: uid, Email: uemail, Username: uusername, Role: urole}, nil
}
