package two_pointers

// 反转字符串中的元音字母: 涉及前后交换的就可以左右双指针往中间指; 注意元音字母包括大小写
func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	bytes := []byte(s)

	for left < right {
		for left < right && !isVowel(bytes[left]) {
			left++
		}
		for left < right && !isVowel(bytes[right]) {
			right--
		}
		if left < right {
			bytes[left], bytes[right] = bytes[right], bytes[left]
			left++
			right--
		}
	}

	return string(bytes)
}

func isVowel(b byte) bool {
	if b == 'a' || b == 'i' || b == 'o' || b == 'e' || b == 'u' || b == 'A' || b == 'I' || b == 'O' || b == 'E' || b == 'U' {
		return true
	}

	return false
}
