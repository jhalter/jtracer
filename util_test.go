package jtracer

import (
	"github.com/google/go-cmp/cmp"
)

// float64Comparer compares the approximate equality of two float64
// https:// pkg.go.dev/github.com/google/go-cmp/cmp#Equal
var float64Comparer = cmp.Comparer(func(a, b float64) bool {
	return (a-b) < epsilon && (b-a) < epsilon
})
