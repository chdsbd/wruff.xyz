package apis

import (
	"errors"
	"net/http"
	"net/url"
	"os"
)

//https://github.com/sjkaliski/go-yo
//Problems with package so I had to modify some stuff

var YO_API = "http://api.justyo.co"

// Yo API Client.
type Client struct {
	Token string
}

// Creates a new Client.
func NewClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) YoUser(username string) error {
	data := url.Values{}
	data.Set("api_token", c.Token)
	data.Set("username", username)
	res, err := http.PostForm(YO_API+"/yo/", data)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("Received response with non 200 status code.")
	}

	return nil
}

func SendYo(username string) error {
	client := NewClient(os.Getenv("yo_apikey"))
	err := client.YoUser(username)
	if err != nil {
		return err
	}
	return nil
}
