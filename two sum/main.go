package twoSum

import "fmt"

/*
Input: nums = [3,2,4], target = 6
Output: [1,2]
*/

func twoSum(nums []int, target int) []int {

	result := []int{}

	for i, left := range nums {
		for j, right := range nums {
			if left+right == target && i != j {
				result = []int{i, j}
				return result
			}
		}
	}

	return result
}

func main() {

	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))

}
