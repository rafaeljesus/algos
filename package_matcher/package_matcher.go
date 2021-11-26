package main

import "fmt"

/*
- M medium box
- L large box

try to fit multiple items into the same box,
but boxes can only only contain one type of product

list of items associated with boxes:
- Cam: 1:M box, 2:L
- Game: 1 doesn't fit M, but 2:L
- Blue: 1:L

Examples:
In: ["Cam"]
Out: [M: ["Cam"]]

In: ["Cam", "Game"]
Out: [M: ["Cam"], L: ["Game"]]

In: ["Game", "Game", "Blue"]
Out: [L: ["Game", "Game"], L: ["Blue"]]
*/
func main() {
	capMap := map[string]map[string]int{
		"Cam": map[string]int{
			"M": 1, "L": 2,
		},
		"Game": map[string]int{
			"L": 2,
		},
		"Blue": map[string]int{
			"L": 1,
		},
	}

	items := []string{"Cam", "Cam", "Game"}
	// out: ["L": ["Cam", "Cam"], "L": ["Game"]]
	itemsWithQty := map[string]int{}
	for _, item := range items {
		itemsWithQty[item]++
	}
	result := []map[string][]string{}
	// [Cam: 1, Game: 1]
	for item, _ := range itemsWithQty {
		if ruleMap, ok := capMap[item]; ok {
			// item = Cam, qty = 2
			// ruleMap = {"M": 1, "L": 2}
			for size, limit := range ruleMap {
				if added > limit {
					continue
				}
				// int and map[string]int
				//quotient := qty / limit
				//remainder := qty % limit
				res := map[string][]string{}
				arr := []string{}
				for i := 0; i < limit; i++ {
					arr = append(arr, item)
				}
				res[size] = arr
				result = append(result, res)
				//qty = remainder
			}
		}
	}
	fmt.Println(result)
}
