package auth

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ValidateUser struct {
	AccessToken string
}

type RefreshToken struct {
	RefreshToken string
}
