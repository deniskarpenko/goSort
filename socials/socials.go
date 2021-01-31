package socials

type Socials interface {
	GetToken() string
	SetToken(token string)
}
