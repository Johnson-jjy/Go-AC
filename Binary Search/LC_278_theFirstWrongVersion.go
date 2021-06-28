package Binary_Search

/**
* Forward declaration of isBadVersion API.
* @param   version   your guess about first bad version
* @return 	 	      true if current version is bad
*			          false if current version is good
*/

// 自己乱写的，只是为了不报错
func isBadVersion(version int) bool {
	return true
}


// ans计数
func firstBadVersion(n int) int {
	left := 1
	right := n
	ans := 1

	for left <= right {
		mid := left + (right - left)/2

		if isBadVersion(mid) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
