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
			name: "A strip pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 0)},
			want: White,
		},
		{
			name: "A strip pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 1, 0)},
			want: White,
		},
		{
			name: "A strip pattern is constant in y",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 2, 0)},
			want: White,
		},
		{
			name: "A strip pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 1)},
			want: White,
		},
		{
			name: "A strip pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 2)},
			want: White,
		},
		{
			name: "A strip pattern is constant in z",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 3)},
			want: White,
		},
		{
			name: "A strip pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0, 0, 0)},
			want: White,
		},
		{
			name: "A strip pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(0.9, 0, 0)},
			want: White,
		},
		{
			name: "A strip pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(1, 0, 0)},
			want: Black,
		},
		{
			name: "A strip pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(-0.1, 0, 0)},
			want: Black,
		},
		{
			name: "A strip pattern alternates in x",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{p: *NewPoint(-1, 0, 0)},
			want: Black,
		},
		{
			name: "A strip pattern alternates in x",
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
				A:         tt.fields.A,
				B:         tt.fields.B,
				Transform: tt.fields.Transform,
			}
			if got := s.ColorAt(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStripePattern_ColorAtObject(t *testing.T) {
	type fields struct {
		A         Color
		B         Color
		Transform Matrix
	}
	type args struct {
		object     Shaper
		worldPoint Tuple
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Color
	}{
		{
			name: "stripes with an object transformation",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: IdentityMatrix,
			},
			args: args{
				object: Sphere{
					Shape: Shape{
						Transform: Scaling(2, 2, 2),
					},
				},
				worldPoint: *NewPoint(1.5, 0, 0),
			},
			want: White,
		},
		{
			name: "stripes with a pattern transformation",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: Scaling(2, 2, 2),
			},
			args: args{
				object:     NewSphere(),
				worldPoint: *NewPoint(1.5, 0, 0),
			},
			want: White,
		},
		{
			name: "stripes with both an object and pattern transformation",
			fields: fields{
				A:         White,
				B:         Black,
				Transform: NewTranslation(0.5, 0, 0),
			},
			args: args{
				object: Sphere{
					Shape: Shape{
						Transform: Scaling(2, 2, 2),
					},
				},
				worldPoint: *NewPoint(1.5, 0, 0),
			},
			want: White,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StripePattern{
				A:         tt.fields.A,
				B:         tt.fields.B,
				Transform: tt.fields.Transform,
			}
			if got := s.ColorAtObject(tt.args.object, tt.args.worldPoint); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAtObject() = %v, want %v", got, tt.want)
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
				A:         tt.fields.A,
				B:         tt.fields.B,
				Transform: tt.fields.Transform,
			}
			if got := s.ColorAt(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ColorAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
