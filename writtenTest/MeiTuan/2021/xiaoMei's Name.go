package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	var n int
	fmt.Scan(&n)
	//var str string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		DigitFlag := false
		finish := false
		str := input.Text()
		//fmt.Println(str)
		if !isLetter(str[0]) {
			fmt.Println("Wrong")
			continue
		}

		for j := 1; j < len(str); j++ {
			if isDigit(str[j]) {
				DigitFlag = true
			} else if !isLetter(str[j]) {
				finish = true
				fmt.Println("Wrong")
				break
			}
		}
		if !finish {
			if DigitFlag {
				fmt.Println("Accept")
			} else {
				fmt.Println("Wrong")
			}
		}
	}
}

func isLetter(b byte) bool {
	if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') {
		return true
	}
	return false
}

func isDigit(b byte) bool {
	//fmt.Printf("%c\n", b)
	if '0' <= b && b <= '9' {
		return true
	}
	return false
}
