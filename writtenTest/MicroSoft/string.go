package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) string {
	// write your code in Go 1.4
	small := make([]int, 26)
	large := make([]int, 26)
	res := byte('a' - 1)

	for i := 0; i < len(S); i++ {
		if 'a' <= S[i] && S[i] <= 'z' {
			small[int(S[i]) - 'a']++
			if large[int(S[i]) - 'a'] > 0 {
				if res < byte('A' + S[i] - 'a') {
					res = byte('A' + S[i] - 'a')
				}
			}
		} else {
			large[int(S[i]) - 'A']++
			if small[int(S[i] - 'A')] > 0 && res < S[i] {
				res = S[i]
			}
		}
	}

	if res == byte('a' - 1) {
		return "NO"
	}
	return string(res)
}
