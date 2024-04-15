package v1

import (
	"net/http"
	"testing"
)

func Test_handlerV1_UploadImageHandler(t *testing.T) {
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
			tt.h.UploadImageHandler(tt.args.w, tt.args.r)
		})
	}
}

func Test_handlerV1_GetImageHandler(t *testing.T) {
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
			tt.h.GetImageHandler(tt.args.w, tt.args.r)
		})
	}
}
