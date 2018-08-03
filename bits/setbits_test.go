package bits

import (
	"testing"
	"math"
)

func Test_countSetBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{name: "success", args: args{5}, wantCount: 2},
		{name: "success2", args: args{1}, wantCount: 1},
		{name: "success3", args: args{0}, wantCount: 0},
		{name: "success_max", args: args{math.MaxInt64}, wantCount: 63},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countSetBits(tt.args.n); gotCount != tt.wantCount {
				t.Errorf("countSetBits() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
