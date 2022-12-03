package jtracer

type Ray struct {
	Origin, Direction Tuple
}

func NewRay(origin, direction Tuple) Ray {
	return Ray{origin, direction}
}

func (r *Ray) Position(t float64) *Tuple {
	return r.Origin.Add(r.Direction.Multiply(t))
}

func (r *Ray) Transform(m Matrix) Ray {
	return Ray{
		*m.MultiplyByTuple(r.Origin),
		*m.MultiplyByTuple(r.Direction),
	}
}
