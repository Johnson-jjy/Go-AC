package Binary_Search

func minEatingSpeed(piles []int, h int) int {
	left := 1
	right := 1000000000

	ans := right
	for left <= right {
		mid := left + (right - left) / 2

		if satisfyH(piles, mid, h) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}

func satisfyH(piles []int, cur int, H int) bool {
	count := 0
	for i := 0; i < len(piles); i++ {
		if piles[i] % cur == 0 {
			count += piles[i] / cur
		} else {
			count += (piles[i] / cur) + 1
		}
		if count > H {
			return false
		}
	}
	return true
}
