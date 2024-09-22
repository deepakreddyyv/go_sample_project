package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readValues(path string) ([]float64, error){
	byteSlice, err := os.ReadFile(path)

	if err != nil {
		return []float64{}, err
	}

	csvData, _ := csv.NewReader(strings.NewReader(string(byteSlice))).ReadAll()

    contents := csvData[1:]

	prices := []float64{}

	for _, val := range contents {
		for _, val2 := range val {
			intConversion, err := strconv.ParseFloat(val2, 64)
            if err != nil {
				return []float64{}, err
			}
            prices = append(prices, intConversion)
		}
	}
	fmt.Println(prices)


	return prices, nil
}

func calculateTaxes(prices []float64, taxRates []float64) [][]float64 {
	var afterTax [][]float64 = [][]float64{}

	for _, tax := range taxRates {
		taxCompute := []float64{}
		for _, prices := range prices {
			res := prices * (1 + (tax / 100))
            taxCompute = append(taxCompute, res)
		}
		afterTax = append(afterTax, taxCompute)
	}
	return afterTax
}

type TaxPricesStore struct {
	TaxRates []float64 `json:"tax_rates"`
	Prices []float64   `json:"prices"`
	AfterTax [][]float64 `json:"after_tax"`
}

func (s *TaxPricesStore) parseTOJson() error {
	data, err := json.Marshal(s)

	if err != nil {
		return errors.New("failed to parse to json")
	}

	stringData := string(data)

	return os.WriteFile("prices_tax.json", []byte(stringData), 0764)
}


func main() {
	var TAXRATES [4]float64 = [4]float64{0, 10, 20, 30}

	prices, _ := readValues("prices.csv")

    afterTax := calculateTaxes(prices, TAXRATES[:])

	fmt.Println(afterTax)

	taxStore := TaxPricesStore {
		TaxRates: TAXRATES[:],
		Prices: prices,
		AfterTax: afterTax,
	}

	err := taxStore.parseTOJson()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfull")

	
}