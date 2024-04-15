package image

import (
	"mime/multipart"
	"testing"
)

func Test_imageInteractor_UploadImage(t *testing.T) {
	type args struct {
		userID   int
		file     multipart.File
		filename string
	}
	tests := []struct {
		name    string
		ii      *imageInteractor
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ii.UploadImage(tt.args.userID, tt.args.file, tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("imageInteractor.UploadImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
