package jtracer

import (
	"reflect"
	"testing"
)

func TestStripePattern_ColorAt(t *testing.T) {
	type fields struct {
		A         Color
		B         Color
		Transform [][]float64
	}
	type args struct {
		p Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name: "A stripe pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 0)},
			want: White,
		},
		{
			name: "A stripe pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 1, 0)},
			want: White,
		},
		{
			name: "A stripe pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 2, 0)},
			want: White,
		},
		{
			name: "A stripe pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 1)},
			want: White,
		},
		{
			name: "A stripe pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 2)},
			want: White,
		},
		{
			name: "A stripe pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 3)},
			want: White,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 0)},
			want: White,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0.9, 0, 0)},
			want: White,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(1, 0, 0)},
			want: Black,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(-0.1, 0, 0)},
			want: Black,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(-1, 0, 0)},
			want: Black,
		},
		{
			name: "A stripe pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(-1.1, 0, 0)},
			want: White,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StripePattern{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			s.SetTransform(tt.fields.Transform)
			if got := s.ColorAt(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckersPattern_ColorAt(t *testing.T) {
	type fields struct {
		A         Color
		B         Color
		Transform Matrix
	}
	type args struct {
		p Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name: "checkers should repeat in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 0, 0),
			},
			want: White,
		},
		{
			name: "checkers should repeat in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0.99, 0, 0),
			},
			want: White,
		},
		{
			name: "checkers should repeat in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(1.01, 0, 0),
			},
			want: Black,
		},
		{
			name: "checkers should repeat in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 0, 0),
			},
			want: White,
		},
		{
			name: "checkers should repeat in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 0.99, 0),
			},
			want: White,
		},
		{
			name: "checkers should repeat in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 1.01, 0),
			},
			want: Black,
		},
		{
			name: "checkers should repeat in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 0, 0.99),
			},
			want: White,
		},
		{
			name: "checkers should repeat in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				p: *NewPoint(0, 0, 1.01),
			},
			want: Black,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := CheckersPattern{
				A: tt.fields.A,
				B: tt.fields.B,
			}
			s.SetTransform(tt.fields.Transform)
			if got := s.ColorAt(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPatternAtShape(t *testing.T) {
	type args struct {
		patterny   Patterny
		shape      Shaper
		worldPoint Tuple
	}
	tests := []struct {
		name string
		args args
		want Color
	}{
		{
			name: "A pattern with an object transformation",
			args: args{
				patterny: NewTestPattern(),
				shape: func() *Sphere {
					s := NewSphere()
					s.SetTransform(Scaling(2, 2, 2))
					s.Material.Pattern = NewTestPattern()
					return s
				}(),
				worldPoint: Tuple{2, 3, 4, 1},
			},
			want: Color{1, 1.5, 2},
		},
		{
			name: "A pattern with a pattern transformation",
			args: args{
				patterny: NewTestPatternWithTransform(Scaling(2, 2, 2)),
				shape: func() *Sphere {
					s := NewSphere()
					s.Material.Pattern = NewTestPatternWithTransform(Scaling(2, 2, 2))
					return s
				}(),
				worldPoint: Tuple{2, 3, 4, 1},
			},
			want: Color{1, 1.5, 2},
		},
		{
			name: "A pattern with both an object and pattern transformation",
			args: args{
				patterny: NewTestPatternWithTransform(NewTranslation(0.5, 1, 1.5)),
				shape: func() *Sphere {
					s := NewSphere()
					s.SetTransform(Scaling(2, 2, 2))
					s.Material.Pattern = NewTestPatternWithTransform(NewTranslation(0.5, 1, 1.5))
					return s
				}(),
				worldPoint: Tuple{2.5, 3, 3.5, 1},
			},
			want: Color{0.75, 0.5, 0.25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PatternAtShape(tt.args.patterny, tt.args.shape, tt.args.worldPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatternAtShape() = %v, want %v", got, tt.want)
			}
		})
	}
}
