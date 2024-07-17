package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req := &http.Request{
		Header: http.Header{
			"Authorization": []string{""},
		},
	}

	got, err := GetAPIKey(req.Header)
	want := "no authorization header included"

	if err != nil {
		want := "malformed authorization header"
		t.Fatalf("expected: %v, got: %v", want, err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
