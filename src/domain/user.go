package domain

type User struct {
	ID    int    `json:"id"`
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

type UserLogin struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
