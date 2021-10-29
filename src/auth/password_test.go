package auth

import "testing"

func TestCreateHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Success case",
			args: args{
				password: "test",
			},
			want:    "548f65b68b529782e3d621955141f75fd907dbe9e485a507ea0a4d75a2710b3505c86f9050e7b7d4c3141a8a54fd91bc72e14366ddefe351a46e56ce78b452df",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateHash(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
