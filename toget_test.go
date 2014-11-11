package toget

import (
	"net/http"
	"testing"
	"time"
)

func TestDo(t *testing.T) {

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := Do(req, 10*time.Second)
	if err != nil {
		t.Error(err)
	}
	if resp == nil {
		t.Error("Expected a response object, got nil")
	}

	resp, err = Do(req, time.Millisecond)
	if err != TimeoutError {
		t.Errorf("Expected %s, Got: %s", TimeoutError, err)
	}
	if resp != nil {
		t.Error("Expected a nil response object, got valid object")
	}

}
