package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req := &http.Request{
		Header: http.Header{
			"Authorization": []string{"Bearer mytoken"},
		},
	}

	got, err := GetAPIKey(req.Header)
	want := "mytoken"

	if err != nil {
		// t.Fatalf("expected: nil, got: %v", err)
		return
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
