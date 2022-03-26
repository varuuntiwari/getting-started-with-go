package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

var (
	nameSlice []name
)

func main() {
	var file string
	// Get filename and open file
	fmt.Print("Enter file name: ")
	fmt.Scanf("%s", &file)
	f, err := os.Open(file)
    if err != nil {
        fmt.Println("Could not read file!")
    }
    defer f.Close()

	out := bufio.NewScanner(f)

	// Read each line
    for out.Scan() {
        arr := (strings.Split(out.Text(), " "))
		// Store name in struct
		tmpName := name{fname: arr[0], lname: arr[1]}
		// Append struct to slice
		nameSlice = append(nameSlice, tmpName)
    }

	// Go through slice and print first and last name
	for _, n := range nameSlice {
		fmt.Printf("First: %v, Last: %v\n", n.fname, n.lname)
	}
}