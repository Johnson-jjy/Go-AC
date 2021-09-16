package main

import "fmt"

func IsNil(i interface{}) bool {
	return i == nil
}

func main()  {
	if IsNil(nil) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
