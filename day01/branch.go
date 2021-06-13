package main

import (
	"fmt"
)

func grade(score int) string {
	var res string
	switch {
	case score < 60:
		res = "C"
	case score < 80:
		res = "B"
	case score <= 100:
		res = "A"
	default:
		panic(fmt.Sprintf("Wrong score : %d", score))
	}
	return res
}

func main() {

	fmt.Println(
		grade(0),
		grade(60),
		grade(100),
		// grade(101),
	)
}
