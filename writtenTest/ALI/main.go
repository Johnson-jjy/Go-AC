package main

import (
	"fmt"
)

type number int

func (n number) print()  {
	fmt.Println(n)
}

func (n *number) pprint()  {
	fmt.Println(*n)
}

func main()  {
	var n number

	defer n.print()
	defer n.pprint()
	defer func() { n.print() }()
	defer func() { n.pprint() }()

	n = 3
}
