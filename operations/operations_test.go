package operations

import (
	"testing"
)

func Test_addition(t *testing.T) {
	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_for_0",
			args: args{-2, 2},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addition(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_substraction(t *testing.T) {
	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_for_minusfour",
			args: args{-2, 2},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := substraction(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("substraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multiplication(t *testing.T) {
	type args struct {
		lhs int
		rhs int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_for_minusfour",
			args: args{-2, 2},
			want: -4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiplication(tt.args.lhs, tt.args.rhs); got != tt.want {
				t.Errorf("multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_division(t *testing.T) {
	type args struct {
		lhs float64
		rhs float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "test_for_one",
			args:    args{1.5, -1.5},
			want:    -1.,
			wantErr: false,
		},
		{
			name:    "test_for_2",
			args:    args{4.222, 2.111},
			want:    2.,
			wantErr: false,
		},
		{
			name:    "test_for_divide_by_zero",
			args:    args{2.623, 0},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := division(tt.args.lhs, tt.args.rhs)
			if (err != nil) != tt.wantErr {
				t.Errorf("division() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("division() = %v, want %v", got, tt.want)
			}
		})
	}
}
