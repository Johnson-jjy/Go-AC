package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution2(S string) int {
	// write your code in Go 1.4
	sum := 0
	for i := 0; i < len(S); i++ {
		sum += int(S[i] - '0')
	}
	res := 0
	if sum % 3 == 0 {
		res -= len(S) - 1
	}

	for i := 0; i < len(S); i++ {
		cur := int(S[i] - '0')
		curSum := sum - cur
		if curSum % 3 == 0 {
			res += 4
		} else {
			res += 3
		}
	}

	return res
}
