package socials

import (
	"io/ioutil"
	"log"
	"net/http"
)

const FACEBOOK_GRAPH_URL = "https://graph.facebook.com/v9.0/"

type Facebook struct {
	token string
}

func (f *Facebook) SetToken(token string) {
	f.token = token;
}

func (f *Facebook) GetToken() string{
	return  f.token;
}

func (f *Facebook) GetMe() (string, bool) {
	if f.token == "" {
		return "", false
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", FACEBOOK_GRAPH_URL + "me/", nil)
	if err != nil {
		return "Bad request", false
	}

	q := req.URL.Query()
	q.Add("access_token", f.GetToken())
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "Errored when sending request to the server", false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body),true
}

func (f *Facebook) GetFollowers(idPage string) (string, bool) {
	if f.token == "" {
		return "", false
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", FACEBOOK_GRAPH_URL + idPage, nil)
	if err != nil {
		return "Bad request", false
	}

	q := req.URL.Query()
	q.Add("access_token", f.GetToken())
	q.Add("fields", "fan_count")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return "Errored when sending request to the server", false
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body),true
}

func NewFacebook() *Facebook {
	return &Facebook{}
}
