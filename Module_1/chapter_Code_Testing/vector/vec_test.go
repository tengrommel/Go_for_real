package vector

import (
	"math"
	"testing"
	"os"
)

func TestNewVector(t *testing.T)  {
	v := New(1,2,4)
	if len(v) != 3 {
		t.Errorf("Expecting vector size %d, but got %d", 3, len(v))
	}
}

func TestVectorString(t *testing.T)  {
	v := New(1,2,3,4)
	if v.String() != "[1,2,3,4]"{
		t.Logf("Expecting [1,2,3,4], but got %s", v.String())
		t.Fail()
	}
}

func TestVectorEqual(t *testing.T)  {
	v1 := New(45, 44, 90)
	var v2 SimpleVector = []float64{45, 44, 90}
	t.Log(v1.Angle(v2), zero, v1.Angle(v2) <= zero)
	if !v1.Eq(v2) {
		t.Logf("Vectors are expected tp be equal")
		t.Fail()
	}
}