package socials

type Facebook struct {
	token string
}

func (f *Facebook) SetToken(token string) {
	f.token = token;
}

func (f *Facebook) GetToken() string{
	return  f.token;
}

func NewFacebook() *Facebook {
	return &Facebook{}
}
