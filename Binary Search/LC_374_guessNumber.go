package Binary_Search

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
*/

// 只是为了不报错
func guess(num int) int{
	return -1
}


func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right - left) / 2

		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == -1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
