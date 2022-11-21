package jtracer

import (
	"reflect"
	"testing"
)

func TestPoint(t *testing.T) {
	p := Tuple{4.3, -4.2, 3.1, 1.0}
	if p.X != 4.3 && p.Y != -4.2 && p.Z != 3.1 && p.W != 1.0 {
		t.Fail()
	}
}

func TestNewPoint(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple
	}{
		{
			name: "creates tuples with w=1",
			args: args{4.3, 4.2, 3.1},
			want: &Tuple{4.3, 4.2, 3.1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewVector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple
	}{
		{
			name: "creates tuples with w=0",
			args: args{
				4.3, 4.2, 3.1,
			},
			want: &Tuple{
				4.3, 4.2, 3.1, 0.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVector(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Add(t1 *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	type args struct {
		a *Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tuple
	}{
		{
			name: "example 1",
			fields: fields{
				3, -2, 5, 1,
			},
			args: args{a: &Tuple{
				-2, 3, 1, 0,
			}},
			want: &Tuple{
				1, 1, 6, 1,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := t.Add(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Subtract(t1 *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	type args struct {
		a *Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tuple
	}{
		{
			name: "subtracting two points",
			fields: fields{
				3, 2, 1, 1.0,
			},
			args: args{
				NewPoint(5, 6, 7),
			},
			want: NewVector(-2, -4, -6),
		},
		{
			name: "subtracting a vector from a point",
			fields: fields{
				3, 2, 1, 1.0,
			},
			args: args{
				NewVector(5, 6, 7),
			},
			want: NewPoint(-2, -4, -6),
		},
		{
			name: "subtracting two vectors",
			fields: fields{
				3, 2, 1, 0.0,
			},
			args: args{
				NewVector(5, 6, 7),
			},
			want: NewVector(-2, -4, -6),
		},
		{
			name:   "subtracting a vector from the zero vector",
			fields: fields{0, 0, 0, 0},
			args: args{
				NewVector(1, -2, 3),
			},
			want: NewVector(-1, 2, -3),
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := t.Subtract(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Negate(t1 *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	tests := []struct {
		name   string
		fields fields
		want   *Tuple
	}{
		{
			name: "negating a tuple",
			fields: fields{
				X: 1,
				Y: -2,
				Z: 3,
				W: -4,
			},
			want: &Tuple{
				X: -1,
				Y: 2,
				Z: -3,
				W: 4,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := t.Negate(); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Negate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Multiply(t1 *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tuple
	}{
		{
			name: "multiplying a tuple by a scalar",
			fields: fields{
				X: 1,
				Y: -2,
				Z: 3,
				W: -4,
			},
			args: args{3.5},
			want: &Tuple{
				X: 3.5,
				Y: -7,
				Z: 10.5,
				W: -14,
			},
		},
		{
			name: "multiplying a tuple by a fraction",
			fields: fields{
				X: 1,
				Y: -2,
				Z: 3,
				W: -4,
			},
			args: args{0.5},
			want: &Tuple{
				X: 0.5,
				Y: -1,
				Z: 1.5,
				W: -2,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := t.Multiply(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTuple_Divide(t1 *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
		W float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tuple
	}{
		{
			name: "multiplying a tuple by a fraction",
			fields: fields{
				X: 1,
				Y: -2,
				Z: 3,
				W: -4,
			},
			args: args{2},
			want: &Tuple{
				X: 0.5,
				Y: -1,
				Z: 1.5,
				W: -2,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Tuple{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := t.Divide(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
