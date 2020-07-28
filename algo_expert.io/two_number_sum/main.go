package main

import "fmt"

func twoNumberSum(array []int, target int) []int {
	// Write your code here.
	winners := []int{}
	for i:=0; i < len(array); i++ {
		for x:=i+1; x < len(array); x++ {
			if array[i] + array[x] == target {
				winners = append(winners, array[i], array[x])
			}
		}
	}
	
	return winners
}
func main () {
	input := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	fmt.Println(twoNumberSum(input, target))
}
