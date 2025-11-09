package user

type User struct {
	ID    string
	Email string

	FirstName string
	LastName  string

	AvatarUrl *string

	HashedPassword string
}
