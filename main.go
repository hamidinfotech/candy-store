package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var filePath string = "customers.csv"

type Customer struct {
	name  string
	candy string
	eaten int
}

type CustomerStat struct {
	Name           string `json:"name"`
	FavouriteSnack string `json:"favouriteSnack"`
	TotalSnacks    int    `json:"totalSnacks"`
}

func main() {
	customers := getAllCustomers()

	customersUniqueCandies := getCustomersUniqueCandies(customers)

	topCustomers := getTopCustomers(customersUniqueCandies)

	jsonData, _ := json.MarshalIndent(topCustomers, "", " ")
	
	fmt.Println(string(jsonData))
}

func getAllCustomers() []Customer {
	data := readFile()

	var customers []Customer

	for i, cutomerData := range data {
		// ignore first csv column
		if i > 0 {
			customers = append(customers, createCustomer(cutomerData))
		}
	}

	return customers
}

func readFile() [][]string {
	file, error := os.Open(filePath)

	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	csv := csv.NewReader(file)

	data, error := csv.ReadAll()

	if error != nil {
		log.Fatal(error)
	}

	return data
}

func createCustomer(customerData []string) Customer {
	eaten, _ := strconv.Atoi(customerData[2])

	var customer Customer
	customer.name = customerData[0]
	customer.candy = customerData[1]
	customer.eaten = eaten

	return customer
}

func getCustomersUniqueCandies(customers []Customer) []Customer {
	customersMap := make(map[string]Customer)

	for _, customer := range customers {
		customerCandy, exist := customersMap[customer.name+customer.candy]

		if exist {
			customerCandy.eaten += customer.eaten
		} else {
			customerCandy = Customer{
				name:  customer.name,
				candy: customer.candy,
				eaten: customer.eaten,
			}
		}

		customersMap[customer.name+customer.candy] = customerCandy
	}

	var customersUniqueCandies []Customer
	for _, customer := range customersMap {
		customersUniqueCandies = append(customersUniqueCandies, customer)
	}

	return customersUniqueCandies
}

func getTopCustomers(customers []Customer) []CustomerStat {
	customersMap := make(map[string]CustomerStat)
	favoriteSnackMap := make(map[string]int)

	for _, customer := range customers {
		customerStat, exist := customersMap[customer.name]

		if exist {
			// Calculate total and favorite snacks
			customerStat.TotalSnacks += customer.eaten

			if customer.eaten > favoriteSnackMap[customer.name] {
				customerStat.FavouriteSnack = customer.candy
				favoriteSnackMap[customer.name] = customer.eaten
			}

		} else {
			customerStat = CustomerStat{
				Name:           customer.name,
				FavouriteSnack: customer.candy,
				TotalSnacks:    customer.eaten,
			}

			favoriteSnackMap[customer.name] = customer.eaten
		}

		customersMap[customer.name] = customerStat
	}

	var topCustomers []CustomerStat
	for _, customer := range customersMap {
		topCustomers = append(topCustomers, customer)
	}

	sort.Slice(topCustomers, func(i, j int) bool { return topCustomers[i].TotalSnacks > topCustomers[j].TotalSnacks })

	return topCustomers
}
