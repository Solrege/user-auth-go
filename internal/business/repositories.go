package business

type UserRepository interface {
	SaveUser(email, password string) (User, error)
}
