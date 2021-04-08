package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var testHandler myHandler
	h := NoSurf(&testHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing -> test successful
	default:
		t.Error(fmt.Sprintf("type is not http.Handler but is %T", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var testHandler myHandler
	h := SessionLoad(&testHandler)

	switch v := h.(type) {
	case http.Handler:
		// do nothing -> test successful
	default:
		t.Error(fmt.Sprintf("type is not http.Handler but is %T", v))
	}
}