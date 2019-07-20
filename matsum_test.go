package matsum_test

import (
	"math"
	"testing"

	"github.com/po3rin/matsum"
	"gonum.org/v1/gonum/mat"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		data mat.Matrix
		f    func(float64) float64
		want float64
	}{
		{
			name: "2*2",
			data: mat.NewDense(2, 2, []float64{1, 2, 3, 4}),
			f:    math.Exp,
			want: 84.7910248837216,
		},
		{
			name: "2*2 with 0",
			data: mat.NewDense(2, 2, []float64{0, 0, 0, 0}),
			f:    math.Exp,
			want: 4,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := matsum.Sum(tt.data, tt.f); got != tt.want {
				t.Fatalf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
