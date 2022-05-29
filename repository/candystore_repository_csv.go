package repository

import (
	"candystore/entity"
	"encoding/csv"
	"log"
	"os"
	"sort"
	"strconv"
)

type CandystoreRepositoryCSV struct {

}

var filePath string = "data/customers.csv"

func (c CandystoreRepositoryCSV) GetTopCustomers() []entity.CustomerStat {
	customers := getAllCustomers()

	customersUniqueCandies := getCustomersUniqueCandies(customers)

	return getTopCustomers(customersUniqueCandies)
}

func getAllCustomers() []entity.Customer {
	data := readFile()

	var customers []entity.Customer

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

func createCustomer(customerData []string) entity.Customer {
	eaten, _ := strconv.Atoi(customerData[2])

	var customer entity.Customer
	customer.Name = customerData[0]
	customer.Candy = customerData[1]
	customer.Eaten = eaten

	return customer
}

func getCustomersUniqueCandies(customers []entity.Customer) []entity.Customer {
	customersMap := make(map[string]entity.Customer)

	for _, customer := range customers {
		customerCandy, exist := customersMap[customer.Name+customer.Candy]

		if exist {
			customerCandy.Eaten += customer.Eaten
		} else {
			customerCandy = entity.Customer{
				Name:  customer.Name,
				Candy: customer.Candy,
				Eaten: customer.Eaten,
			}
		}

		customersMap[customer.Name+customer.Candy] = customerCandy
	}

	var customersUniqueCandies []entity.Customer
	for _, customer := range customersMap {
		customersUniqueCandies = append(customersUniqueCandies, customer)
	}

	return customersUniqueCandies
}

func getTopCustomers(customers []entity.Customer) []entity.CustomerStat {
	customersMap := make(map[string]entity.CustomerStat)
	favoriteSnackMap := make(map[string]int)

	for _, customer := range customers {
		customerStat, exist := customersMap[customer.Name]

		if exist {
			// Calculate total and favorite snacks
			customerStat.TotalSnacks += customer.Eaten

			if customer.Eaten > favoriteSnackMap[customer.Name] {
				customerStat.FavouriteSnack = customer.Candy
				favoriteSnackMap[customer.Name] = customer.Eaten
			}

		} else {
			customerStat = entity.CustomerStat{
				Name:           customer.Name,
				FavouriteSnack: customer.Candy,
				TotalSnacks:    customer.Eaten,
			}

			favoriteSnackMap[customer.Name] = customer.Eaten
		}

		customersMap[customer.Name] = customerStat
	}

	var topCustomers []entity.CustomerStat
	for _, customer := range customersMap {
		topCustomers = append(topCustomers, customer)
	}

	sort.Slice(topCustomers, func(i, j int) bool { return topCustomers[i].TotalSnacks > topCustomers[j].TotalSnacks })

	return topCustomers
}