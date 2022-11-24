package jtracer

import (
	"math"
	"reflect"
	"testing"
)

func TestNewTranslation(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "creates a new translation matrix",
			args: args{
				x: 1,
				y: 2,
				z: 3,
			},
			want: Matrix{
				{1.0, 0, 0, 1},
				{0, 1.0, 0, 2},
				{0, 0, 1.0, 3},
				{0, 0, 0, 1.0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTranslation(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTranslation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransformationsAreAppliedInSequence(t *testing.T) {
	p := NewPoint(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := NewTranslation(10, 5, 7)

	p2 := a.MultiplyByTuple(*p)
	if !p2.Equals(NewPoint(1, -1, 0)) {
		t.Fail()
	}

	p3 := b.MultiplyByTuple(p2)
	if !p3.Equals(NewPoint(5, -5, 0)) {
		t.Fail()
	}

	p4 := c.MultiplyByTuple(p3)
	if !p4.Equals(NewPoint(15, 0, 7)) {
		t.Fail()
	}
}

func TestViewTransform(t *testing.T) {
	type args struct {
		from *Tuple
		to   *Tuple
		up   *Tuple
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "the transformation matrix for the default orientation",
			args: args{
				from: NewPoint(0, 0, 0),
				to:   NewPoint(0, 0, -1),
				up:   NewVector(0, 1, 0),
			},
			want: IdentityMatrix,
		},
		{
			name: "a view transformation matrix looking in positive z direction",
			args: args{
				from: NewPoint(0, 0, 0),
				to:   NewPoint(0, 0, 1),
				up:   NewVector(0, 1, 0),
			},
			want: Scaling(-1, 1, -1),
		},
		{
			name: "the view transformation moves the world",
			args: args{
				from: NewPoint(0, 0, 8),
				to:   NewPoint(0, 0, 0),
				up:   NewVector(0, 1, 0),
			},
			want: NewTranslation(0, 0, -8),
		},
		{
			name: "an arbitrary view transformation",
			args: args{
				from: NewPoint(1, 3, 2),
				to:   NewPoint(4, -2, 8),
				up:   NewVector(1, 1, 0),
			},
			want: Matrix{
				{-0.50709, 0.50709, 0.67612, -2.36643},
				{0.76772, 0.60609, 0.12122, -2.82843},
				{-0.35857, 0.59761, -0.71714, 0.00000},
				{0.00000, 0.00000, 0.00000, 1.00000},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ViewTransform(tt.args.from, tt.args.to, tt.args.up); !got.Equal(tt.want) {
				t.Errorf("ViewTransform() = %v, want %v", got, tt.want)
			}
		})
	}
}
