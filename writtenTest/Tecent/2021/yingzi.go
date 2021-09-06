package main

import (
	"fmt"
	"sort"
)

func main()  {
	// 求质数
	primes := make([]int, 0)
	isPrimes := make([]bool, 100001)
	for i := 2; i <= 100000; i++ {
		isPrimes[i] = true
	}
	for i := 2; i <= 100000; i++ {
		if isPrimes[i] {
			primes = append(primes, i)
			for j := i + i; j <= 100000; j = j + i {
				isPrimes[j] = false
			}
		}
	}
	//fmt.Println(primes)

	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	b := make([]int, n)

	var cur int
	check := make([]int, 100001)

	for i := 0; i < n; i++ {
		a[i] = 1
		fmt.Scan(&cur)
		//fmt.Println("?", cur)
		if check[cur] > 0 {
			a[i] = check[cur]
			continue
		}
		need := cur
		cnt := make([]int, len(primes))
		index := 0
		for cur > 1 {
			for cur % primes[index] == 0 {
				cur /= primes[index]
				cnt[index]++
			}
			a[i] *= cnt[index] + 1
			index++
		}
		check[need] = a[i]
		//fmt.Println("1",cur, check[cur])
	}

	for i := 0; i < n; i++ {
		b[i] = 1
		fmt.Scan(&cur)
		if check[cur] > 0 {
			b[i] = check[cur]
			continue
		}
		need := cur
		cnt := make([]int, len(primes))
		index := 0
		for cur > 1 {
			for cur % primes[index] == 0 {
				cur /= primes[index]
				cnt[index]++
			}
			b[i] *= cnt[index] + 1
			index++
		}

		check[need] = b[i]
		//fmt.Println("2",cur, check[cur])
	}

	sort.Ints(a)
	sort.Ints(b)

	//fmt.Println(a)
	//fmt.Println(b)

	res := 0
	left := 0
	right := 0
	for left < n {
		if a[left] > b[right] {
			res++
			left++
			right++
		} else {
			left++
		}
	}

	fmt.Println(res)
}

