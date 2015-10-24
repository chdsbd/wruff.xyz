package apis

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestSendTweet(t *testing.T) {
	username := os.Getenv("twitter_TestUsername")
	message := fmt.Sprintf("Testing %v", time.Now())
	err := SendTweet(username, message)
	if err != nil {
		t.Error(err)
	}
}
