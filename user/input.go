package user

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required,max=255"`
	Occupation string `json:"occupation" binding:"required,max=255"`
	Email      string `json:"email" binding:"required,email,max=255"`
	Password   string `json:"password"  binding:"required,min=6"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
