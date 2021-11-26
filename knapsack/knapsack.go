package main

func Knapsack(profits, weights []int, capacity int) int {
	return knapsack(profits, weights, capacity, 0)
}

func knapsack(profits, weights []int, capacity, currIdx int) int {
	if capacity <= 0 || currIdx >= len(profits) {
		return 0
	}
	var profit1, profit2 int
	if weights[currIdx] <= capacity {
		profit1 = profits[currIdx] + knapsack(profits, weights, capacity-weights[currIdx], currIdx+1)
		// 1 + 6 + 10
	}
	profit2 = knapsack(profits, weights, capacity, currIdx+1)
	return max(profit1, profit2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//
// profits = [1, 6, 10, 16]
// weights = [1, 2, 3, 5]

// capacity = 7
// currIdx = 0
// weights[currIds] <= capacity
// 1 <= 7? y
// profits[currIdx]
// 1
// capacity - weights[currIdx]
// 7 - 1 = 6
// currIdx = 1
// #follow-up profit1 = 1 + recurstion

// capacity = 6
// currIdx >= len(profits)? n
// 1 >= 4
// weights[currIdx] <= capacity?
// 2 <= 6?y
// profits[currIdx]
// 6
// #follow-up profit1 = 6 + recursion
// capacity - weights[currIdx]
// 6 - 2 = 4

// capacity = 4
// currIdx = 2
// currIdx >= len(profits)?
// 2 >= 4? n
// weights[currIdx]
// 3 <= 4? y
// profits[currIdx]
// 10
// #follow-up profit1 = 10 + recursion
// capacity - weights[currIdx]
// 4 - 2 = 2

// capacity = 2
// currIdx = 3
// weights[currIdx] <= capacity?
// 5 <= 2? n
// max 10, 0
