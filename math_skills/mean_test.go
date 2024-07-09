package math_skills

import "testing"

func TestAverage(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test1",
			args: args{numbers: []float64{1, 2, 3, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.numbers); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}
