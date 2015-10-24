package apis

import "testing"

func TestSendEmail(t *testing.T) {
    err := SendEmail("mail@example.com", "Testing")
    if err != nil {
        t.Error(err)
    }
}
