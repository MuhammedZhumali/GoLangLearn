package main

import "fmt"

func main() {
	//Task 1
	fmt.Println(countWords([]string{"go", "java", "go", "python", "go", "java"}))

	//Task 2
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

	//Task 3
	fmt.Println(firstRepeated([]int{5, 1, 3, 1, 5}))

	//Task 4
	fmt.Println(twoSumMap([]int{2, 7, 11, 15}, 9))
}

func countWords(words []string) map[string]int {
	count := map[string]int{}

	for i := 0; i < len(words); i++ {
		count[words[i]]++
	}
	return count
}

func containsDuplicate(nums []int) bool {
	dup := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if dup[nums[i]] {
			return true
		}

		dup[nums[i]] = true
	}
	return false
}

func firstRepeated(nums []int) int {
	rep := map[int]bool{}
	var n int

	for i := 0; i < len(nums); i++ {
		if rep[nums[i]] {
			n = nums[i]
			return n
		}

		rep[nums[i]] = true
	}
	return -1
}

func twoSumMap(nums []int, target int) []int {
	seen := map[int]int{}

	for i := 0; i < len(nums); i++ {
		current := nums[i]
		need := target - current

		index, ok := seen[need]
		if ok {
			return []int{index, i}
		}

		seen[current] = i
	}

	return []int{}
}
