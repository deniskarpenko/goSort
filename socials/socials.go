package socials

type Socials interface {
	GetToken() string
	SetToken(token string)
	GetFollowers(idPage string) (string, bool)
	GetMe() (string, bool)
}

func FactorySocial(keySocial string) Socials {
	switch keySocial {
	case "fb":
		return NewFacebook()
	case "instagram":
		return NewInstagram()
	}
	return &Facebook{}
}
