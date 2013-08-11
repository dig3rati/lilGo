package lilGo

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestShorten(t *testing.T) {
	t.Error("No Shorten method defined yet " + randomURL(30))
}

func TestExpand(t *testing.T) {
	t.Error("No Expand method defined yet " + randomURL(10))
}

func randomURL(length int) (randomURL string) {
	randomCode := make([]byte, 3*length)
	for i := 1; i <= 3; i++ {
		b := make([]byte, length)
		rand.Read(b)
		en := base64.URLEncoding
		d := make([]byte, en.EncodedLen(len(b)))
		en.Encode(d, b)
		randomCode += d
	}
	return fmt.Sprintf("http://lil.go/%s", randomCode)
}
