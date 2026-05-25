package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	fmt.Print(isValid("(s)"))

}

func isValid(s string) bool {
	count := 0

	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count++
		}

		if s[i] == ')' {
			count--
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

func twoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}
	return res
}
