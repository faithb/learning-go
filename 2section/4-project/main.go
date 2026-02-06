package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64{
	"T-SHIRT": 12.00,
	"MUG":     11.50,
	"HAT":     10.00,
	"BOOT":    5.05,
}

func calculateItemPrice(itemPrice string) (float64, bool) {
	basePrice, found := productPrices[itemPrice]

	if !found {
		if strings.HasSuffix(itemPrice, "_SALE") {
			originalItemCode := strings.TrimSuffix(itemPrice, "_SALE")
			basePrice, found = productPrices[originalItemCode]
			if found {
				salePrice := basePrice * 0.9
				fmt.Printf("--- Item %s (Sale! Original: %.2f, salePrice: %.2f)\n", originalItemCode, basePrice, salePrice)
			}
		}

		fmt.Printf("--- Item: %s (Product not found)\n", itemPrice)

		return 0.0, false

	}

	return basePrice, true
}

func main() {
	fmt.Println("---- simple sale order processing ----")

	saleOrder := []string{
		"T-SHIRT",
		"MUG_SALE",
		"HAT",
		"BOOT_SALE",
		"JACKET",
	}

	var totalPrice float64
	for _, item := range saleOrder {
		price, found := calculateItemPrice(item)
		if found {
			totalPrice += price
		}
	}

	fmt.Printf("Total price: %.2f\n", totalPrice)
}
