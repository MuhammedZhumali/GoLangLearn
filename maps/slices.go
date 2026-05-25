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

	//Task 6
	fmt.Println(contains(nums, 6))

	//Task 7
	fmt.Println(countTarget([]int{1, 2, 2, 3, 2}, 2))

	//Task 8
	fmt.Println(filterEven(nums))

	//Task 9
	fmt.Println(removeLast(nums))

	//Task 10
	stack := []int{}
	stack = append(stack, 10)
	stack = append(stack, 20)
	stack = append(stack, 30)
	top := stack[len(stack)-1]
	fmt.Println("top:", top)

	stack = stack[:len(stack)-1]
	fmt.Println("stack:", stack)
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

func contains(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return true
		}
	}
	return false
}

func countTarget(nums []int, target int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

func filterEven(nums []int) []int {
	var filter []int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			filter = append(filter, nums[i])
		}
	}
	return filter
}

func removeLast(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	nums = nums[:len(nums)-1]
	return nums
}
