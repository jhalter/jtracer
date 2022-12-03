package jtracer

import (
	"reflect"
	"testing"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *Tuple
		direction *Tuple
	}
	tests := []struct {
		name string
		args args
		want Ray
	}{
		{
			name: "creating and querying a ray",
			args: args{
				origin:    NewPoint(1, 2, 3),
				direction: NewVector(4, 5, 6),
			},
			want: Ray{
				Origin:    &Tuple{1, 2, 3, 1},
				Direction: &Tuple{4, 5, 6, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRay(tt.args.origin, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_Position(t *testing.T) {
	type fields struct {
		Origin    *Tuple
		Direction *Tuple
	}
	type args struct {
		t float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tuple
	}{
		{
			name: "computing a point from a distance",
			fields: fields{
				Origin:    NewPoint(2, 3, 4),
				Direction: NewVector(1, 0, 0),
			},
			args: args{0},
			want: NewPoint(2, 3, 4),
		},
		{
			name: "computing a point from a distance",
			fields: fields{
				Origin:    NewPoint(2, 3, 4),
				Direction: NewVector(1, 0, 0),
			},
			args: args{1},
			want: NewPoint(3, 3, 4),
		},
		{
			name: "computing a point from a distance",
			fields: fields{
				Origin:    NewPoint(2, 3, 4),
				Direction: NewVector(1, 0, 0),
			},
			args: args{-1},
			want: NewPoint(1, 3, 4),
		},
		{
			name: "computing a point from a distance",
			fields: fields{
				Origin:    NewPoint(2, 3, 4),
				Direction: NewVector(1, 0, 0),
			},
			args: args{2.5},
			want: NewPoint(4.5, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			if got := r.Position(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Position() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRay_Transform(t *testing.T) {
	type fields struct {
		Origin    *Tuple
		Direction *Tuple
	}
	type args struct {
		m Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Ray
	}{
		{
			name: "translating a ray",
			fields: fields{
				Origin:    NewPoint(1, 2, 3),
				Direction: NewVector(0, 1, 0),
			},
			args: args{m: NewTranslation(3, 4, 5)},
			want: Ray{
				Origin:    NewPoint(4, 6, 8),
				Direction: NewVector(0, 1, 0),
			},
		},
		{
			name: "scaling a ray",
			fields: fields{
				Origin:    NewPoint(1, 2, 3),
				Direction: NewVector(0, 1, 0),
			},
			args: args{m: Scaling(2, 3, 4)},
			want: Ray{
				Origin:    NewPoint(2, 6, 12),
				Direction: NewVector(0, 3, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Ray{
				Origin:    tt.fields.Origin,
				Direction: tt.fields.Direction,
			}
			if got := r.Transform(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Transform() = %v, want %v", got, tt.want)
			}
		})
	}
}
