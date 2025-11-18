package main

import "fmt"

func main() {
	cakeCost := 0.99
	cakeSalesTaxRate := 0.075
	milkCost := 2.75
	milkSalesTaxRate := 0.015
	butterCost := 0.87
	butterSalesTaxRate := 0.02

	cakeTotalSalesTax := cakeCost * cakeSalesTaxRate
	milkTotalSalesTax := milkCost * milkSalesTaxRate
	butterTotalSalesTax := butterCost * butterSalesTaxRate

	fmt.Println("Sales Tax Total: ", cakeTotalSalesTax+milkTotalSalesTax+butterTotalSalesTax)
}
