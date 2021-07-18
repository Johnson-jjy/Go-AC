package main

import "fmt"

func main(){
	nums := [3]int{1,2,3}
	nums[1], nums[2] = nums[2], nums[1]
	fmt.Printf("%v", nums)
}
