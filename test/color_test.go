package test

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	a := 0.7
	r := ((1 - a) * 0) + a*255
	g := ((1 - a) * 0) + a*255
	b := ((1 - a) * 0) + a*255
	fmt.Println(r, g, b)

	a = 0.1
	r = ((1 - a) * 255) + a*0
	g = ((1 - a) * 255) + a*200
	b = ((1 - a) * 255) + a*220
	fmt.Println(r, g, b)
}
