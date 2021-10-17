package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Book struct {
	Name  string
	Price float64
}

type Books []*Book

func (s Books) Len() int      { return len(s) }
func (s Books) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ByName struct{ Books }

func (s ByName) Less(i, j int) bool { return s.Books[i].Name < s.Books[j].Name }

type ByPrice struct{ Books }

func (s ByPrice) Less(i, j int) bool { return s.Books[i].Price < s.Books[j].Price }

func main() {
	fmt.Println("sortProductMapExample()")
	sortProductMapExample()
	fmt.Println("sortProductArrayStructExample()")
	sortProductArrayStructExample()
	//fmt.Println("sortWrapperExample()")
	//sortWrapperExample()
	//fmt.Println("sortMapExample()")
	//sortMapExample()
}

type Product struct {
	Name          string
	Price, Rating float64
}

type Products []*Product

// input = [
// 	[Macbook, 1.149, 8.9],
// 	[Lenovo, 1.399,, 8.3],
// 	[Dell, 1.009, 8.9],
// ]
// output:
//  Dell - 1.009 - 8.9
//  Macbook - 1.249 - 8.9
//  Mlenovo - 1.399 - 8.3
func sortProductArrayStructExample() {
	products := []*Product{
		{"Macbook", 1.149, 8.9},
		{"Lenovo", 1.399, 8.3},
		{"Dell", 1.099, 8.9},
	}
	// sort by rating and price
	sort.Slice(products, func(i, j int) bool {
		if products[i].Rating == products[j].Rating {
			return products[i].Price < products[j].Price
		}
		return products[i].Rating > products[j].Rating
	})
	for _, p := range products {
		fmt.Printf("Name: %s, Rating: %.2f, Price: %.4f\n", p.Name, p.Rating, p.Price)
	}
}

func sortProductMapExample() {
	productsMap := map[string][]float64{
		"Macbook": []float64{1.149, 8.9},
		"Lenovo":  []float64{1.399, 8.3},
		"Dell":    []float64{1.099, 8.9},
	}
	// create a array of ratings
	var ratings []string
	ratingsMap := make(map[string]float64, len(productsMap))
	for name, arr := range productsMap {
		// name:rating = price
		// Macbook:8.9 = 1.149
		// Lenovo:8.3 = 1.399
		// Dell:8.9 = 1.099
		nameWithRating := fmt.Sprintf("%s:%f", name, arr[1])
		ratingsMap[nameWithRating] = arr[0]
		ratings = append(ratings, nameWithRating)
	}
	// order by rating and price
	sort.Slice(ratings, func(i, j int) bool {
		// order desc by rating and price
		irating := strings.Split(ratings[i], ":")[1]
		jrating := strings.Split(ratings[j], ":")[1]
		if irating == jrating {
			iprice, ok := ratingsMap[ratings[i]]
			if !ok {
				panic(fmt.Sprintf("rating[%v] should be a key in ratingsMap", ratings[i]))
			}
			jprice, ok := ratingsMap[ratings[j]]
			if !ok {
				panic(fmt.Sprintf("rating[%v] should be a key in ratingsMap", ratings[j]))
			}
			return iprice < jprice
		}
		// order desc by rating
		return irating > jrating
	})

	for _, nameWithRating := range ratings {
		name := strings.Split(nameWithRating, ":")[0]
		rating := strings.Split(nameWithRating, ":")[1]
		if price, ok := ratingsMap[nameWithRating]; ok {
			ratingParsed, err := strconv.ParseFloat(rating, 64)
			if err != nil {
				panic(fmt.Sprintf("unable to parse rating to float64: %v", err))
			}
			fmt.Printf("Name: %s, Rating: %.2f, Price: %.4f\n", name, ratingParsed, price)
		}
	}
}

func sortMapExample() {
	booksByName := map[string]float64{
		"Golang":     45.0,
		"Unix":       33.9,
		"Linux":      27.0,
		"Algorithms": 47.0,
		"SRE":        37.0,
	}

	var keys []string
	for k, _ := range booksByName {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	fmt.Println("Sorted asc by name\n")
	for _, name := range keys {
		if price, ok := booksByName[name]; ok {
			fmt.Printf("Name: %v, Price: %v\n", name, price)
		}
	}

	fmt.Println("Sorted desc by price\n")
	booksByPrice := map[float64]string{
		45.0: "Golang",
		33.9: "Unix",
		27.0: "Linux",
		47.0: "Algorithms",
		37.0: "SRE",
	}
	var prices []float64
	for p, _ := range booksByPrice {
		prices = append(prices, p)
	}
	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})
	for _, price := range prices {
		if name, ok := booksByPrice[price]; ok {
			fmt.Printf("Name: %v, Price: %v\n", name, price)
		}
	}
}

func sortWrapperExample() {
	books := []*Book{
		{"Golang", 45.0},
		{"Unix", 33.9},
		{"Linux", 27.0},
		{"Algorithms", 47.0},
		{"SRE", 37.0},
	}

	fmt.Println("Sorting desc books ByName")
	sort.Slice(books, func(i, j int) bool {
		return books[i].Name > books[j].Name
	})
	printBooks(books)
	fmt.Println("Sorting desc books ByName")

	fmt.Println("Sorting asc books ByName")
	sort.Slice(books, func(i, j int) bool {
		return books[i].Name < books[j].Name
	})
	printBooks(books)
	fmt.Println("Sorting asc books ByName")

	// or use this form
	fmt.Println("Sorting books ByName")
	sort.Sort(ByName{books})
	printBooks(books)
	fmt.Println("---------------------------")

	fmt.Println("Sorting books ByPrice")
	sort.Sort(ByPrice{books})
	printBooks(books)
	fmt.Println("---------------------------")

	fmt.Println("Sorting books ByPrice desc")
	sort.Sort(sort.Reverse(ByPrice{books}))
	printBooks(books)
	fmt.Println("---------------------------")
}

func printBooks(books []*Book) {
	for _, b := range books {
		fmt.Printf("name: %s, price: %f\n", b.Name, b.Price)
	}
}
