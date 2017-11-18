package vector

import (
	"bytes"
	"math"
	"strconv"
)

const (
	zero = 1.0e-7
)

type Vector interface {
	String() string
	Eq(other Vector) bool
	Add(other Vector) Vector
	Sub(other Vector) Vector
	Scale(factor float64)
	DotProd(other Vector) float64
	Angle(other Vector) float64
	Mag() float64
	Unit() Vector
}

type SimpleVector []float64

func New(elems ...float64) SimpleVector {
	return SimpleVector(elems)
}