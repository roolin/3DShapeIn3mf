package meshGen

import (
	"fmt"
	"testing"
)

func Test_segmentLength(t *testing.T) {
	x := segmentLength(30, 45)
	fmt.Println(x)
	if x != 22.961005941905388 {
		t.Error()
	}
}

func Test_degreeToRadian(t *testing.T) {
	x := degreeToRadian(45)
	fmt.Println(x)
	if x != 0.7853981633974483 {
		t.Error()
	}
}

func Test_XYShift(t *testing.T) {
	x, y := XYShift(30, 45)
	fmt.Printf("x: %f, y: %f\n", x, y)
	x, y = XYShift(30, 90)
	fmt.Printf("x: %f, y: %f\n", x, y)
	t.Error()
}
