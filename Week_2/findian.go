package main

import (
	"strings"
	"regexp"
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input);
	result, err := regexp.MatchString("^i.*a.*n$", strings.ToLower(input))
	if err != nil {
		fmt.Println(err)
	}
	if result {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}