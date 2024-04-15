package chat

import (
	"chat2/internal/entities"
	"reflect"
	"testing"
)

func Test_chatInteractor_LlamaChatRequest(t *testing.T) {
	type args struct {
		requestData *entities.ChatReq
	}
	tests := []struct {
		name    string
		ci      *chatInteractor
		args    args
		want    *entities.LlamaResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ci.LlamaChatRequest(tt.args.requestData)
			if (err != nil) != tt.wantErr {
				t.Errorf("chatInteractor.LlamaChatRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chatInteractor.LlamaChatRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
