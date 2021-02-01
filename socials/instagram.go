package socials

type Instagram struct {
	token string
}

func (i *Instagram) SetToken(token string) {
	i.token = token;
}

func (i *Instagram) GetToken() string{
	return  i.token;
}

func (i *Instagram) GetFollowers(idPage string) (string, bool) {
	return "followers", true
}

func (i *Instagram) GetMe() (string, bool) {
	return "me", true
}

func NewInstagram() *Instagram {
	return &Instagram{}
}