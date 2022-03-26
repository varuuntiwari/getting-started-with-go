package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var details = make(map[string]string)
	var tmp string
	inp := bufio.NewReader(os.Stdin)
	tmp, _ = inp.ReadString('\n')
	details["name"] = tmp
	tmp, _ = inp.ReadString('\n')
	details["address"] = tmp
	obj, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(obj))
	}
}