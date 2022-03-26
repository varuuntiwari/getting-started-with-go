package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := make([]int64, 3)
	var i int64
	for true {
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			return
		}
		slice = append(slice, i)
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] < slice[j]
		})
		for _, x := range slice {
			fmt.Print(x, " ")
		}
		fmt.Println()
	}
}
