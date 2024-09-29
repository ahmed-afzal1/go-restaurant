package requests

type RegisterRequest struct {
	Email                string `json:"email" form:"email" binding:"required"`
	Phone                string `json:"phone" form:"phone"`
	Password             string `json:"password" form:"password" binding:"required,min=8"`
	PasswordConfirmation string `json:"password_confirmation" form:"password_confirmation" binding:"eqfield=Password"`
}
