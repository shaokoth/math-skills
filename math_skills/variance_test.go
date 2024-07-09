package math_skills

import "testing"

func TestVariance(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test3",
			args: args{numbers: []float64{1, 2, 3, 4, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.numbers); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}
