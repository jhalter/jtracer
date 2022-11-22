package jtracer

import (
	"reflect"
	"testing"
)

func TestColor_Add(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		a *Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "adding colors",
			fields: fields{
				Red:   0.9,
				Green: 0.6,
				Blue:  0.75,
			},
			args: args{a: &Color{
				Red:   0.7,
				Green: 0.1,
				Blue:  0.25,
			}},
			want: &Color{
				Red:   1.6,
				Green: 0.7,
				Blue:  1.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				Red:   tt.fields.Red,
				Green: tt.fields.Green,
				Blue:  tt.fields.Blue,
			}

			if got := c.Add(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Subtract(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		a *Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "subtracting colors",
			fields: fields{
				Red:   0.9,
				Green: 0.6,
				Blue:  0.75,
			},
			args: args{a: &Color{
				Red:   0.7,
				Green: 0.1,
				Blue:  0.25,
			}},
			want: &Color{
				Red:   0.2,
				Green: 0.5,
				Blue:  0.5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				Red:   tt.fields.Red,
				Green: tt.fields.Green,
				Blue:  tt.fields.Blue,
			}
			if got := c.Subtract(tt.args.a); !got.Equals(tt.want) {
				t.Errorf("Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_MultiplyByScalar(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		a float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "multiplying a color by a scalar",
			fields: fields{
				0.2, 0.3, 0.4,
			},
			args: args{2.0},
			want: &Color{
				0.4, 0.6, 0.8,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				Red:   tt.fields.Red,
				Green: tt.fields.Green,
				Blue:  tt.fields.Blue,
			}
			if got := c.MultiplyByScalar(tt.args.a); !got.Equals(tt.want) {
				t.Errorf("MultiplyByScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColor_Multiply(t *testing.T) {
	type fields struct {
		Red   float64
		Green float64
		Blue  float64
	}
	type args struct {
		a *Color
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Color
	}{
		{
			name: "multiplying colors",
			fields: fields{
				1.0, 0.2, 0.4,
			},
			args: args{
				&Color{0.9, 1, 0.1},
			},
			want: &Color{0.9, 0.2, 0.04},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				Red:   tt.fields.Red,
				Green: tt.fields.Green,
				Blue:  tt.fields.Blue,
			}
			if got := c.Multiply(tt.args.a); !got.Equals(tt.want) {
				t.Errorf("Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
