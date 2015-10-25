package apis

import (
	"os"
	"testing"
)

func TestSendYo(t *testing.T) {
	err := SendYo(os.Getenv("yo_username"))
	if err != nil {
		t.Error(err)
	}
}
