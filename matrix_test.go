package jtracer

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"math"
	"reflect"
	"testing"
)

func TestMatrix_Multiply(t *testing.T) {
	type args struct {
		m2 Matrix
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Matrix
	}{
		{
			name: "multiplying two matrices",
			m: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			args: args{
				m2: Matrix{
					{-2, 1, 2, 3},
					{3, 2, 1, -1},
					{4, 3, 6, 5},
					{1, 2, 7, 8},
				},
			},
			want: Matrix{
				{20, 22, 50, 48},
				{44, 54, 114, 108},
				{40, 58, 110, 102},
				{16, 26, 46, 42},
			},
		},
		{
			name: "multiplying a matrix by the identity matrix",
			m: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			args: args{
				m2: IdentityMatrix,
			},
			want: Matrix{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Multiply(tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_MultiplyByTuple(t *testing.T) {
	type args struct {
		t Tuple
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want *Tuple
	}{
		{
			name: "a matrix multiplied by a tuple",
			m: Matrix{
				{1, 2, 3, 4},
				{2, 4, 4, 2},
				{8, 6, 4, 1},
				{0, 0, 0, 1},
			},
			args: args{
				Tuple{1, 2, 3, 1},
			},
			want: &Tuple{18, 24, 33, 1},
		},
		{
			name: "the identity matrix multiplied by a tuple",
			m:    IdentityMatrix,
			args: args{
				Tuple{1, 2, 3, 1},
			},
			want: &Tuple{1, 2, 3, 1},
		},
		{
			name: "multiplying a point by a translation matrix",
			m:    NewTranslation(5, -3, 2),
			args: args{
				*NewPoint(-3, 4, 5),
			},
			want: &Tuple{2, 1, 7, 1},
		},
		{
			name: "multiplying a point by the inverse of a translation matrix",
			m: func() Matrix {
				m := NewTranslation(5, -3, 2)
				return m.Inverse()
			}(),
			args: args{
				*NewPoint(-3, 4, 5),
			},
			want: &Tuple{-8, 7, 3, 1},
		},
		{
			name: "translation does not affect vectors",
			m:    NewTranslation(5, -3, 2),
			args: args{
				*NewVector(-3, 4, 5),
			},
			want: &Tuple{-3, 4, 5, 0},
		},
		{
			name: "a scaling matrix applied to a point",
			m:    Scaling(2, 3, 4),
			args: args{
				*NewPoint(-4, 6, 8),
			},
			want: &Tuple{-8, 18, 32, 1},
		},
		{
			name: "a scaling matrix applied to a vector",
			m:    Scaling(2, 3, 4),
			args: args{
				*NewVector(-4, 6, 8),
			},
			want: &Tuple{-8, 18, 32, 0},
		},
		{
			name: "multiplying by  by the inverse of a scaling matrix",
			m: func() Matrix {
				m := Scaling(2, 3, 4)
				return m.Inverse()
			}(),
			args: args{
				*NewVector(-4, 6, 8),
			},
			want: &Tuple{-2, 2, 2, 0},
		},
		{
			name: "reflection is scaling by a negative value",
			m:    Scaling(-1, 1, 1),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: &Tuple{-2, 3, 4, 1},
		},
		{
			name: "rotating a point half quarter around the x-axis",
			m:    RotationX(math.Pi / 4),
			args: args{
				*NewPoint(0, 1, 0),
			},
			want: NewPoint(0, math.Sqrt2/2, math.Sqrt2/2),
		},
		{
			name: "rotating a point full quarter around the x-axis",
			m:    RotationX(math.Pi / 2),
			args: args{
				*NewPoint(0, 1, 0),
			},
			want: NewPoint(0, 0, 1),
		},
		{
			name: "the inverse of an x-rotation rotates in the opposite direction",
			m: func() Matrix {
				m := RotationX(math.Pi / 4)
				return m.Inverse()
			}(),
			args: args{
				*NewPoint(0, 1, 0),
			},
			want: NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2),
		},
		{
			name: "rotating a point half quarter around the y-axis",
			m:    RotationY(math.Pi / 4),
			args: args{
				*NewPoint(0, 0, 1),
			},
			want: NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2),
		},
		{
			name: "rotating a point full quarter around the y-axis",
			m:    RotationY(math.Pi / 2),
			args: args{
				*NewPoint(0, 0, 1),
			},
			want: NewPoint(1, 0, 0),
		},
		{
			name: "rotating a point half quarter around the z-axis",
			m:    RotationZ(math.Pi / 4),
			args: args{
				*NewPoint(0, 1, 0),
			},
			want: NewPoint(-math.Sqrt2/2, math.Sqrt2/2, 0),
		},
		{
			name: "rotating a point full quarter around the z-axis",
			m:    RotationZ(math.Pi / 2),
			args: args{
				*NewPoint(0, 1, 0),
			},
			want: NewPoint(-1, 0, 0),
		},
		{
			name: "a shearing transformation moves x in proportion to y",
			m:    Shearing(1, 0, 0, 0, 0, 0),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(5, 3, 4),
		},
		{
			name: "a shearing transformation moves x in proportion to z",
			m:    Shearing(0, 1, 0, 0, 0, 0),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(6, 3, 4),
		},
		{
			name: "a shearing transformation moves y in proportion to x",
			m:    Shearing(0, 0, 1, 0, 0, 0),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(2, 5, 4),
		},
		{
			name: "a shearing transformation moves y in proportion to z",
			m:    Shearing(0, 0, 0, 1, 0, 0),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(2, 7, 4),
		},
		{
			name: "a shearing transformation moves z in proportion to x",
			m:    Shearing(0, 0, 0, 0, 1, 0),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(2, 3, 6),
		},
		{
			name: "a shearing transformation moves z in proportion to y",
			m:    Shearing(0, 0, 0, 0, 0, 1),
			args: args{
				*NewPoint(2, 3, 4),
			},
			want: NewPoint(2, 3, 7),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.MultiplyByTuple(tt.args.t); !cmp.Equal(got, tt.want, float64Comparer) {
				t.Errorf("MultiplyByTuple() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Determinant(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want float64
	}{
		{
			name: "calculating the determinant of a 2x2 matrix",
			m: Matrix{
				{1, 5},
				{-3, 2},
			},
			want: 17,
		},
		{
			name: "calculating the determinant of a 3x3 matrix",
			m: Matrix{
				{1, 2, 6},
				{-5, 8, -4},
				{2, 6, 4},
			},
			want: -196,
		},
		{
			name: "calculating the determinant of a 4x4 matrix",
			m: Matrix{
				{-2, -8, 3, 5},
				{-3, 1, 7, 3},
				{1, 2, -9, 6},
				{-6, 7, 7, -9},
			},
			want: -4071,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Determinant(); got != tt.want {
				t.Errorf("Determinant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Submatrix(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want Matrix
	}{
		{
			name: "a submatrix of a 3x3 matrix is a 2x2 matrix",
			m: Matrix{
				{1, 5, 1},
				{-3, 2, 7},
				{0, 6, -3},
			},
			args: args{0, 2},
			want: Matrix{
				{-3, 2},
				{0, 6},
			},
		},
		{
			name: "a submatrix of a 4x4 matrix is a 3x3 matrix",
			m: Matrix{
				{-6, 1, 1, 6},
				{-8, 5, 8, 6},
				{-1, 0, 8, 2},
				{-7, 1, -1, 1},
			},
			args: args{2, 1},
			want: Matrix{
				{-6, 1, 6},
				{-8, 8, 6},
				{-7, -1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("hai")
			if got := tt.m.Submatrix(tt.args.row, tt.args.col); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Submatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Minor(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want float64
	}{
		{
			name: "calculating the minor of a 3x3 matrix",
			m: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			args: args{
				row: 1,
				col: 0,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Minor(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("Minor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Cofactor(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		m    Matrix
		args args
		want float64
	}{
		{
			name: "calculating the cofactor of a 3x3 matrix for row 0, col 0",
			m: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			args: args{
				row: 0,
				col: 0,
			},
			want: -12,
		},
		{
			name: "calculating the cofactor of a 3x3 matrix for row 1, col 0",
			m: Matrix{
				{3, 5, 0},
				{2, -1, -7},
				{6, -1, 5},
			},
			args: args{
				row: 1,
				col: 0,
			},
			want: -25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Cofactor(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("Cofactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Inverse(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Matrix
	}{
		{
			name: "calculating the inverse of a matrix",
			m: Matrix{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			},
			want: Matrix{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			},
		},
		{
			name: "calculating the inverse of another matrix",
			m: Matrix{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			},
			want: Matrix{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			},
		},
		{
			name: "calculating the inverse of a third matrix",
			m: Matrix{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			},
			want: Matrix{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Inverse(); !cmp.Equal(got, tt.want, float64Comparer) {
				t.Errorf("Inverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMatrix_Transpose(t *testing.T) {
	tests := []struct {
		name string
		m    Matrix
		want Matrix
	}{
		{
			name: "transposing a matrix",
			m: Matrix{
				{0, 9, 3, 0},
				{9, 8, 0, 8},
				{1, 8, 5, 3},
				{0, 0, 5, 8},
			},
			want: Matrix{
				{0, 9, 1, 0},
				{9, 8, 8, 0},
				{3, 0, 5, 5},
				{0, 8, 3, 8},
			},
		},
		{
			name: "transposing a matrix",
			m:    IdentityMatrix,
			want: IdentityMatrix,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Transpose(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
