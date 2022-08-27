package server

import (
	"reflect"
	"testing"
)

func TestNewServer(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Checking type of NewServer",
			want: "*server.Server",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(); !reflect.DeepEqual(
				reflect.TypeOf(got).String(),
				tt.want,
			) {
				t.Errorf("NewServer() = %v, want %v", reflect.TypeOf(got).String(), tt.want)
			}
		})
	}
}
