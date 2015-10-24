package apis

import "testing"

func TestSendYo(t *testing.T) {
	err := SendYo(os.Env("yo_username"))
	if err != nil {
		t.Error(err)
	}
}
