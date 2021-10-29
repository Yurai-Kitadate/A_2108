package area

import (
	"reflect"
	"testing"
)

func Test_getAPIResponse(t *testing.T) {
	type args struct {
		pref int
	}
	tests := []struct {
		name    string
		args    args
		want    response
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAPIResponse(tt.args.pref)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAPIResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAPIResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
