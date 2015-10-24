// Copyright 2011 Arne Roomann-Kurrik
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apis

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
)

func loadCredentials() (client *twittergo.Client) {
	config := &oauth1a.ClientConfig{
		ConsumerKey:    os.Getenv("twitter_ConsumerKey"),
		ConsumerSecret: os.Getenv("twitter_ConsumerSecret"),
	}
	user := oauth1a.NewAuthorizedConfig(
		os.Getenv("twitter_AccessToken"),
		os.Getenv("twitter_AcessTokenSecret"),
	)
	client = twittergo.NewClient(config, user)
	return
}

func SendTweet(username, message string) error {
	var (
		err    error
		client *twittergo.Client
		req    *http.Request
		resp   *twittergo.APIResponse
		tweet  *twittergo.Tweet
	)
	client = loadCredentials()
	data := url.Values{}
	data.Set("status", fmt.Sprintf("@%s %s", username, message)) //TODO: INSERT STUFF HERE!!!
	body := strings.NewReader(data.Encode())
	req, err = http.NewRequest("POST", "/1.1/statuses/update.json", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err = client.SendRequest(req)
	if err != nil {
		return err
	}
	tweet = &twittergo.Tweet{}
	err = resp.Parse(tweet)
	if err != nil {
		return err
	}
	return nil
}
