package main

import (
	"errors"
	"fmt"
)

func main() {
	for {
		const USD_TO_EUR float64 = 0.88
		const USD_TO_RUB float64 = 81.15
		const EUR_TO_RUB float64 = USD_TO_RUB / USD_TO_EUR
		amount, err := UserInput()
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Printf("the amount is: %v \n", amount)
		original_currency, target_currency, err := UserInputCurrency()
		if err != nil {
			fmt.Print(err)
			continue
		} else {
			fmt.Printf("conversion amount of %v from %s to %s is: ", amount, original_currency, target_currency)
		}

		switch {
		case original_currency == "USD" && target_currency == "EUR":
			fmt.Println(amount * USD_TO_EUR)
		case original_currency == "USD" && target_currency == "RUB":
			fmt.Println(amount * USD_TO_RUB)
		case original_currency == "EUR" && target_currency == "USD":
			fmt.Println(amount / USD_TO_EUR)
		case original_currency == "EUR" && target_currency == "RUB":
			fmt.Println(amount * EUR_TO_RUB)
		case original_currency == "RUB" && target_currency == "USD":
			fmt.Println(amount / USD_TO_RUB)
		case original_currency == "RUB" && target_currency == "EUR":
			fmt.Println(amount / EUR_TO_RUB)
		default:
			fmt.Printf("unsupported currency pair: %s to %s", original_currency, target_currency)
		}
	}

}

func UserInput() (float64, error) {
	var amount float64
	fmt.Println("please enter the amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		return 0, errors.New("enter a value above 0")

	} else {
		return amount, nil
	}

}

func UserInputCurrency() (string, string, error) {
	var original_currency string
	var target_currency string
	fmt.Println("please choose original currency from EUR, RUB , USD")
	fmt.Scan(&original_currency)
	fmt.Println("please choose the target currency from EUR RUB USD")
	fmt.Scan(&target_currency)

	if (original_currency == "EUR" || original_currency == "USD" || original_currency == "RUB") && (target_currency == "EUR" || target_currency == "USD" || target_currency == "RUB") {
		return original_currency, target_currency, nil
	} else {
		return "", "", errors.New("Please enter either EUR USD OR RUB")
	}

}
