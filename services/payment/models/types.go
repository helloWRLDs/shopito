package models

import "time"

type ReceiptData struct {
	CompanyName string    `json:"company_name,omitempty"`
	Items       []Item    `json:"items"`
	Customer    string    `json:"customer"`
	Date        time.Time `json:"date,omitempty"`
}

type Cart struct {
	Items  []Item `json:"items"`
	CVV    string `json:"cvv"`
	Name   string `json:"name"`
	Number string `json:"number"`
	Email  string `json:"email"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
