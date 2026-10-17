package main

import "fmt"

func main() {
	a := []int{1, 2, 8, 9, 10}
	target := 10
	seen := map[int]int{}

	for idx, val := range a {
		delta := target - val
		seen_idx, found := seen[delta]
		if found {
			fmt.Printf("indexes: %v", []int{idx, seen_idx})
			break
		} else {
			seen[val] = idx
		}
		fmt.Printf("idx: %v, val: %v, delta: %v, seen: %v\n", idx, val, delta, seen)
	}

}
