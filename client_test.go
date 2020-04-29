package dtrobot

import "testing"

func TestPost(t *testing.T) {
	Post("http://xxx.com", []byte("test"), WithHeader("ok", "json"))
}
