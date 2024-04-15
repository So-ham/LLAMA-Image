package image

import (
	"reflect"
	"testing"
)

func Test_imageInteractor_GetImage(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		ii      *imageInteractor
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ii.GetImage(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageInteractor.GetImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageInteractor.GetImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
