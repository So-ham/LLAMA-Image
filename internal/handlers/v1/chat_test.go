package v1

import (
	"net/http"
	"testing"
)

func Test_handlerV1_ChatHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		h    *handlerV1
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.ChatHandler(tt.args.w, tt.args.r)
		})
	}
}
