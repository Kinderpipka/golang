package main

import (
	"fmt"

	"github.com/Kinderpipka/golang/internal/app"
)

func main() {

	a := []string{"s", "f", "g"}
	x := "sasd"
	// app.Ykazatel()
	// app.Radius()
	// app.Nedel()
	fmt.Println(app.Contains(a, x))
}
