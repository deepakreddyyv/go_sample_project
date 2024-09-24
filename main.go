package main

import (
	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/tax"
)
func main() {

    var TAXRATES [4]float64 = [4]float64{0, 10, 20, 30}

    fm := &filemanager.FileManager{
		ReadFilepath: "prices.csv",
		WriteFilePath: "prices_tax.json",
	}


	ts := tax.TaxStore{
		IoManager: fm,
		TaxRates: TAXRATES[:],
		Prices: []float64{},
		AfterTax: [][]float64{},
	}

	ts.TransformValues()

	err := ts.CalculateTax()

	if err != nil {
		panic(err)
	}
}