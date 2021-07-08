package logics

import "fmt"

func FindLongConsecutive() {
	var input = []string{"W", "L", "W", "W", "W", "W", "L", "W", "W", "L", "L", "L", "W", "W", "L", "L"}

	curr := 0
	max := 0

	for _, itr := range input {
		if itr == "L" && curr > max {
			max = curr
			curr = 0
		} else if itr == "L" {
			curr = 0
		} else {
			curr++
		}

	}
	fmt.Print("w")
	if max > curr {
		fmt.Println(max)
	} else {
		fmt.Println(curr)
	}
}
