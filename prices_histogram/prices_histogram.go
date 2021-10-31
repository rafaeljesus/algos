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
	bucketsMap := map[int]string{
		0:  "",
		10: "",
		20: "",
		30: "",
		40: "",
		50: "",
		60: "",
		70: "",
		80: "",
		90: "",
	}
	for _, price := range prices {
		for bucket, s := range bucketsMap {
			if price == bucket {
				bucketsMap[bucket] = s + "*"
			} else if price > bucket && price < bucket+10 {
				bucketsMap[bucket] = s + "*"
			}
		}
	}
	buckets := make([]int, 0)
	for bucket, _ := range bucketsMap {
		buckets = append(buckets, bucket)
	}
	sort.Slice(buckets, func(i, j int) bool {
		return buckets[i] < buckets[j]
	})
	for _, orderedBucket := range buckets {
		if histo, ok := bucketsMap[orderedBucket]; ok {
			fmt.Printf("%v: %s\n", orderedBucket, histo)
		}
	}
}
