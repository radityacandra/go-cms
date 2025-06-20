package types

type RegisterUserInput struct {
	Username string
	Password string
	FullName string
}

type RegisterUserOutput struct {
	Id string
}

type FindUserByUsernameInput struct {
	Username string `db:"username"`
}

type FindUserByUsernameOutput struct {
	Id       string
	Username string
	Password string
}

type DetailUserOutput struct {
	Id       string
	Username string
	Acls     []string
}
