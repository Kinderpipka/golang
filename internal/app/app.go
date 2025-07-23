package app

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

func GetMax(a ...int) {
	sum := 0
	for _, item := range a {
		if item > sum {
			sum = item
		}
	}
	fmt.Println(sum)

}

func Sozdanie() {
	chitka, err := os.Open("cmd/main/06_task_in.txt")
	if err != nil {
		panic(err)
	}
	defer chitka.Close()

	scanenrr := bufio.NewScanner(chitka)
	lineCount := 0
	zapis, err := os.Create("cmd/main/out.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		zapis.Close()
		fmt.Printf("%d", lineCount)
	}()

	lineNumber := 1
	zap := bufio.NewWriter(zapis)
	for scanenrr.Scan() {
		zap.WriteString(fmt.Sprintf("%d,%s \n", lineNumber, scanenrr.Text()))
		lineNumber++
		lineCount++
	}
	zap.Flush()
	defer logTime()

}

func logTime() {
	time1 := time.Now()
	fmt.Println(time.Since(time1))

}
