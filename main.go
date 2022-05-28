package main

import (
	"encoding/json"
	"fmt"
	"candystore/candystore"
)

func main() {
	topCustomers := candystore.TopCustomers()
	
	jsonData, _ := json.MarshalIndent(topCustomers, "", " ")
	
	fmt.Println(string(jsonData))
}
