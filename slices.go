package main

import "fmt"

func main() {
	//Task 1
	nums := []int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	//Task 2
	fmt.Println(sum(nums))

	//Task 3
	fmt.Println(findMax(nums))

	//Task 4
	fmt.Println(countEven(nums))

	//Task 5
	fmt.Println(revers(nums))
}

func sum(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

func findMax(nums []int) int {
	var max int = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func countEven(nums []int) int {
	var count int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			count++
		}
	}
	return count
}

func revers(nums []int) []int {
	var rev []int
	for i := len(nums) - 1; i >= 0; i-- { //for rever i = len(arr)-1
		rev = append(rev, nums[i])
	}
	return rev
}
