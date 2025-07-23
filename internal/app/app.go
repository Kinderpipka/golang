package app

import (
	"fmt"
)

func Ykazatel() {
	var a *int
	b := 23
	a = &b
	fmt.Println(*a)
	*a = 50
	fmt.Println(b)
}

func Radius() {
	var r *float64
	dlinna := float64(35)
	result := dlinna / (2 * 3.14)
	r = &result
	fmt.Println(*r)

	ploshad := 3.14 * *r
	fmt.Println(ploshad)
}
