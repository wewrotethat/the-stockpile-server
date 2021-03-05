type User struct {
	ID       uint `json:"id"`
	UserName string `json:"userName"`
	FullName string  `json:"fullName"`
	Email    string `json:"email"`
	Phone    string  `json:"phone"`
	Password string  `json:"-"`
}

