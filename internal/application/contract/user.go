package contract

type CreateUserContract struct {
	Email    string
	Password string
}

type CreateUserResponse struct {
	ID    string
	Email string
}
