package randomizer

import (
	"reflect"
	"testing"
)

func TestBytes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Check random length", args{n: 15}, 15, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Bytes(tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("Bytes() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		nBytes int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"Test case", args{nBytes: 3}, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := String(tt.args.nBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("String() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("String() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestToken(t *testing.T) {
	tests := []struct {
		name    string
		want    int
		wantErr bool
	}{
		{"Test length", 44, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Token()
			if (err != nil) != tt.wantErr {
				t.Errorf("Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Token() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
