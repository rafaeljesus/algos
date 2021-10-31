package main

import (
	"fmt"
	"sort"
)

/*
prices = [10, 11, 13, 22, 24, 38]
bucket = {0: "", 10: "", 20: "", 30: "", 40: "", 50: "", 60: "", 70: "", 80: "", 90: ""}
Expected Output:
0:
10: ***
20: **
30: *
40:
50:
60:
70:
80:
90:
*/

func main() {
	prices := []int{10, 11, 13, 22, 24, 38}
	bucketMap := map[int]string{0: "", 10: "", 20: "", 30: "", 40: "", 50: "", 60: "", 70: "", 80: "", 90: ""}
	// O(n * m)
	for _, price := range prices {
		for bucket, str := range bucketMap {
			if price >= bucket && price < bucket+10 {
				bucketMap[bucket] = fmt.Sprintf("%s*", str)
			}
		}
	}
	bucketsRange := make([]int, 0)
	for bucket, _ := range bucketMap {
		bucketsRange = append(bucketsRange, bucket)
	}
	sort.Slice(bucketsRange, func(i, j int) bool {
		return bucketsRange[i] < bucketsRange[j]
	})
	for _, orderedBucket := range bucketsRange {
		if str, ok := bucketMap[orderedBucket]; ok {
			fmt.Printf("%d: %s\n", orderedBucket, str)
		}
	}
}
