package auth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success case",
			args: args{
				id: 110,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			jwt, err := VerifyToken(got)
			if err != nil {
				t.Errorf("GenerateToken() = %v, Verify failed with %v", got, err)
			}
			fmt.Printf("%#v\n", jwt)
		})
	}
}
