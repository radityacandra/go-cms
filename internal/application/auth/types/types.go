package types

type LoginInput struct {
	Username string
	Password string
}

type LoginOutput struct {
	Token     string
	ExpiredAt int64
}
