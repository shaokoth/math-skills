package math_skills

import "testing"

func TestStdDeviation(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test2",
			args: args{numbers: []float64{1, 2, 3, 4, 5}},
			want: 1.4142135623730951,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StdDeviation(tt.args.numbers); got != tt.want {
				t.Errorf("StdDeviation() = %v, want %v", got, tt.want)
			}
		})
	}
}
