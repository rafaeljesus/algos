package main

import (
	"fmt"
	"math/rand"
	"time"
)

func SecretSantaPartOne(people []string) map[string]string {
	selection := make(map[string]string)
	t := time.Now().UnixNano()
	rand.Seed(t)
	for len(people) > 0 {
		current := people[0]
		secretIdx := rand.Intn(len(people))
		secret := people[secretIdx]
		_, secretAlreadySelected := selection[secret]
		for current == secret || secretAlreadySelected {
			secretIdx = rand.Intn(len(people))
			secret = people[secretIdx]
			_, secretAlreadySelected = selection[secret]
			if current == secret || secretAlreadySelected {
				continue
			}
			break
		}
		selection[current] = secret
		people = append(people[:secretIdx], people[secretIdx+1:]...)
		people = people[1:]
	}
	return selection
}

func SecretSantaPartTwo(people []string, history map[int]int) map[string]string {
	selection := make(map[string]string)
	t := time.Now()
	rand.Seed(t.UnixNano())
	var currentIdx int
	for len(people) > 0 {
		current := people[0]
		secretIdx := rand.Intn(len(people))
		secret := people[secretIdx]
		_, secretAlreadySelected := selection[secret]
		pairAlreadySelectedInTheLast3Years := findInHistory(t, history, currentIdx, secretIdx)
		for current == secret || secretAlreadySelected || pairAlreadySelectedInTheLast3Years {
			secretIdx = rand.Intn(len(people))
			secret = people[secretIdx]
			_, secretAlreadySelected = selection[secret]
			pairAlreadySelectedInTheLast3Years = findInHistory(t, history, currentIdx, secretIdx)
			if current == secret || secretAlreadySelected || pairAlreadySelectedInTheLast3Years {
				continue
			}
			break
		}
		selection[current] = secret
		people = append(people[:secretIdx], people[secretIdx+1:]...)
		people = people[1:]
		currentIdx++
	}
	return selection
}

func findInHistory(t time.Time, history map[int]int, currentIdx, secretIdx int) bool {
	if currentIdx == secretIdx {
		return false
	}
	for i := 0; i < 3; i++ {
		t = t.AddDate(-i, 0, 0)
		if pairIdxSum, ok := history[t.Year()]; ok {
			if currentIdx+secretIdx == pairIdxSum {
				fmt.Printf("currentIdx: %d, secretIdx: %d, pairIdxSum: %d\n", currentIdx, secretIdx, pairIdxSum)
				return true
			}
		}
	}
	return false
}

func main() {
	people := []string{"Raf", "Ju", "Sophia", "Pamela", "David", "Junior"}
	out := SecretSantaPartOne(people)
	fmt.Printf("partOne: %v\n", out)

	t := time.Now()
	history := map[int]int{
		t.AddDate(-1, 0, 0).Year(): 0 + 2, // 2020: Raf + Sophia
		t.AddDate(-2, 0, 0).Year(): 1 + 3, // 2019: Ju + Pamela
		t.AddDate(-3, 0, 0).Year(): 4 + 5, // 2018: David + Junior
	}
	people = []string{"Raf", "Ju", "Sophia", "Pamela", "David", "Junior"}
	out = SecretSantaPartTwo(people, history)
	fmt.Printf("partTwo: %v\n", out)
}
