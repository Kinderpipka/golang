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

func Nedel() {

	n := []string{"понедельник", "вторник", "среда", "четверг", "пятница", "суббота", "воскресенье"}
	b := n[:5]
	n = n[5:]

	sum := append(b, n...)

	fmt.Println(b)
	fmt.Println(n)
	fmt.Println(sum)

}

func Contains(a []string, x string) bool {
	for _, item := range a {
		if item == x {
			return true
		}

	}
	return false
}
