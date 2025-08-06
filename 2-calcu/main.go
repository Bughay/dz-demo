package main

import (
	"errors"
	"fmt"
)

var exchangeRates = map[string]map[string]float64{
	"USD": {
		"EUR": 0.88,
		"RUB": 81.15,
		"USD": 1,
	},
	"EUR": {
		"USD": 1 / 0.88,
		"RUB": 81.15 / 0.88,
		"EUR": 1,
	},
	"RUB": {
		"USD": 1 / 81.15,
		"EUR": 0.88 / 81.15,
		"RUB": 1,
	},
}

func main() {
	for {
		amount, err := UserInput()
		if err != nil {
			fmt.Print(err)
			continue
		}
		fmt.Printf("The amount is: %v\n", amount)

		originalCurrency, targetCurrency, err := UserInputCurrency()
		if err != nil {
			fmt.Print(err)
			continue
		}

		fmt.Printf("Conversion amount of %v from %s to %s is: ", amount, originalCurrency, targetCurrency)

		rate, ok := exchangeRates[originalCurrency][targetCurrency]
		if !ok {
			fmt.Printf("unsupported currency pair: %s to %s\n", originalCurrency, targetCurrency)
			continue
		}

		fmt.Println(amount * rate)
	}
}

func UserInput() (float64, error) {
	var amount float64
	fmt.Println("Please enter the amount:")
	_, err := fmt.Scan(&amount)
	if err != nil {
		return 0, errors.New("invalid input")
	}

	if amount <= 0 {
		return 0, errors.New("enter a value above 0\n")
	}
	return amount, nil
}

func UserInputCurrency() (string, string, error) {
	var originalCurrency, targetCurrency string

	fmt.Println("Please choose original currency from EUR, RUB, USD:")
	_, err := fmt.Scan(&originalCurrency)
	if err != nil {
		return "", "", errors.New("invalid input")
	}

	fmt.Println("Please choose the target currency from EUR, RUB, USD:")
	_, err = fmt.Scan(&targetCurrency)
	if err != nil {
		return "", "", errors.New("invalid input")
	}

	if !isValidCurrency(originalCurrency) || !isValidCurrency(targetCurrency) {
		return "", "", errors.New("please enter either EUR, USD or RUB\n")
	}

	return originalCurrency, targetCurrency, nil
}

func isValidCurrency(currency string) bool {
	_, exists := exchangeRates[currency]
	return exists
}
