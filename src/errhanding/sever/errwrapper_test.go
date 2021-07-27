package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(1)
}

func TestErrWrapper(t *testing.T) {
	tests := []struct {
		h    appHandler
		code int
		msg  string
	}{
		{errPanic, 500, ""},
	}

	//for _,tt := range tests {
	//	wrapper := errWrapper(tt.h)
	//	resp := httptest.NewRecorder()
	//	req := http.NewRequest(http.MethodGet, "", nil)
	//
	//
	//}
}
