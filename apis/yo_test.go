package apis

import "testing"

func TestSendYo(t *testing.T) {
	err := SendYo("CHDSBD")
	if err != nil {
		t.Error(err)
	}
}
