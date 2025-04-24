package models

type Car struct {
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	VIN   string `json:"vin"`
	Owner string `json:"owner"`
}