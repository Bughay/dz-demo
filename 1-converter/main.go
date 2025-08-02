package main

import "fmt"

func main() {
	const USD_TO_EUR float64 = 0.88
	const USD_TO_RUB float64 = 81.15
	const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
	amount := UserInput()
	fmt.Printf("the amount is: %v \n", amount)
}

func UserInput() (amount float64) {
	fmt.Println("please enter the amount: ")
	fmt.Scan(&amount)
	return amount
}

func Calculate(amount float64, original_currency string, target_currency string) {}
