package robot

import "fmt"

// Given a robot which can only move in four directions, UP(U), DOWN(D), LEFT(L), RIGHT(R).
// Given a string consisting of instructions to move. Output the coordinates of a robot after executing the instructions.
// Initial position of robot is at origin(0, 0).
// Approach:
// 	For x-coordinate, increase the count for Right move and decrease count for a left move.
// 	For y-coordinate, increase the count for the up move and down count for a left move.

// Input : move = “UDDLRL”
// Output : (-1, -1)
func RobotHandleMove(move string) []int {
	var xAxis, yAxis int
	for _, m := range move {
		switch string(m) {
		case "U":
			yAxis++
		case "D":
			yAxis--
		case "R":
			xAxis++
		case "L":
			xAxis--
		default:
			panic(fmt.Sprintf("unexpected move: %v", m))
		}
	}
	return []int{xAxis, yAxis}
}

// Approach:
// Count number of up movements (U), down movements (D), left movements (L) and right movements (R) as countUp, countDown, countLeft and countRight respectively.
// Final x-coordinate will be: (countRight – countLeft) and y-coordinate will be (countUp – countDown).
func RobotHandleMove2(move string) []int {
	var countUp, countDown, countLeft, countRight int
	for _, m := range move {
		switch string(m) {
		case "U":
			countUp++
		case "D":
			countDown++
		case "R":
			countRight++
		case "L":
			countLeft++
		default:
			panic(fmt.Sprintf("unexpected move: %v", m))
		}
	}
	return []int{countRight - countLeft, countUp - countDown}
}

// Input: UDDLRL
// Expected: [-1, -1]

// countUp = 0
// countDown = 0
// countRight = 0
// countLeft = 0
// m = U
// countUp = 1

// countUp = 1
// countDown = 0
// countRight = 0
// countLeft = 0
// m = D
// countDown = 1

// countUp = 1
// countDown = 1
// countRight = 0
// countLeft = 0
// m = D
// countDown = 2

// countUp = 1
// countDown = 2
// countRight = 0
// countLeft = 0
// m = L
// countLeft = 1

// countUp = 1
// countDown = 2
// countRight = 0
// countLeft = 1
// m = R
// countRight = 1

// countUp = 1
// countDown = 2
// countRight = 1
// countLeft = 1
// m = L
// countLeft = 2

// up = 1
// down = 2
// right  = 1
// left = 2

// right - left +(1 - 2) = -1
// up - down +(1 - 2) = -1
