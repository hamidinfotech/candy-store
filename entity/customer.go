package entity

type Customer struct {
	Name string `json:"name"`
	Candy string `json:"candy"`
	Eaten int `json:"eaten"`
}