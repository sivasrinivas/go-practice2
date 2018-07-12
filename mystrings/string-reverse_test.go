package mystrings

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"success", args{"sample test"}, "tset elpmas"},
		{"empty", args{""}, ""},
		{"blank", args{"    "}, "    "},
		{"palindrome", args{"liril"}, "liril"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReverseString(tt.args.s); got != tt.want {
				t.Errorf("reverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}
