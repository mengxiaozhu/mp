package mp

type Mp struct {
	Token string
}

func NewMp(token string) (mp *Mp) {
	return &Mp{Token: token}
}
