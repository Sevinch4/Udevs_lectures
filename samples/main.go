package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"os"
)

func main() {
	//products read file
	file, err := os.Open("data/products.json")
	if err != nil {
		fmt.Println("error is while opening file", err.Error())
	}
	defer file.Close()
	products := []Product{}

	if err := json.NewDecoder(file).Decode(&products); err != nil {
		panic(err)
	}
	//fmt.Println(products)
	//braches.json
	branch, err := os.Open("data/branches.json")
	if err != nil {
		fmt.Println("error is while opening file", err.Error())
	}
	defer file.Close()
	braches := []Branch{}

	if err := json.NewDecoder(branch).Decode(&braches); err != nil {
		panic(err)
	}
	//fmt.Println(braches)

	brachMap := make(map[int]string)
	for _, b := range braches {
		brachMap[b.ID] = b.Name
	}
	//transaction.json
	transact, err := os.Open("data/branch_pr_transaction.json")
	if err != nil {
		fmt.Println("error is while opening file", err.Error())
	}
	defer file.Close()
	transaction := []Transaction{}

	if err := json.NewDecoder(transact).Decode(&transaction); err != nil {
		panic(err)
	}
	transactionMap := make(map[int]int)
	for _, t := range transaction {
		transactionMap[t.BranchID]++
	}
	type Count struct {
		Key   int
		Value int
	}
	var counts = []Count{}
	for key, value := range transactionMap {
		counts = append(counts, Count{key, value})
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Value > counts[j].Value
	})
	for _, v := range counts {
		fmt.Println("Branch name: ", brachMap[v.Key], "->", v.Value)
	}
	productMap := make(map[int]int)
	for _, p := range products {
		productMap[p.ID] = p.Price
	}
	sum := make(map[int]int)
	for _, t := range transaction {
		sum[t.BranchID] += productMap[t.ProuctID] * t.Quantity
	}
	var counts2 = []Count{}
	for key, value := range sum {
		counts2 = append(counts2, Count{key, value})
	}
	sort.Slice(counts2, func(i, j int) bool {
		return counts2[i].Value > counts2[j].Value
	})
	for _, v := range counts2 {
		fmt.Println("Branch name: ", brachMap[v.Key], "->", v.Value)
	}
	// for key, value := range sum {
	// 	fmt.Println(brachMap[key], value)
	// }
	//task-3
	pMap := make(map[int]string)
	for _, p := range products {
		pMap[p.ID] = p.Name
	}
	countMap := make(map[int]int)
	for _, t := range transaction {
		countMap[t.ProuctID]++
	}
	var counts3 = []Count{}
	for key, value := range transactionMap {
		counts3 = append(counts3, Count{key, value})
	}
	sort.Slice(counts3, func(i, j int) bool {
		return counts3[i].Value > counts3[j].Value
	})
	for _, v := range counts3 {
		fmt.Println("Product name: ", pMap[v.Key], "->", v.Value)
	}
	//fmt.Println(transactionMap)
	//branch_products.json
	//brach_pr,err := os.

}

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Price       int        `json:"price"`
	Category_id int        `json:"category_id"`
	Category    Categories `json:"categories"`
	Braches     []Branch   `json:"braches"`
}
type Categories struct {
	ID   int
	Name string
}
type Branch struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
type BranchProducts struct {
	BranchID     int           `json:"branch_id"`
	ProuctID     int           `json:"product_id"`
	Quantity     int           `json:"quantity"`
	Transactions []Transaction `json:"transaction"`
}
type Transaction struct {
	ID        int    `json:"id"`
	BranchID  int    `json:"branch_id"`
	ProuctID  int    `json:"product_id"`
	Type      string `json:"type"`
	Quantity  int    `json:"quantity"`
	CreatedAt string `json:"create_at"`
}
