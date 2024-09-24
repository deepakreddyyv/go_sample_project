package tests

import (
	"reflect"
	"testing"

	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/tax"
)




func TestCalculateTax(t *testing.T) {

	fm := &filemanager.FileManager{
		ReadFilepath: "prices.csv",
		WriteFilePath: "prices_tax.json",
	}

	ts := &tax.TaxStore{
		IoManager: fm,
		TaxRates: []float64{1, 2, 3, 4, 5},
		Prices: []float64{10, 20, 30, 40, 50, 60},
		AfterTax: [][]float64{},
	}

	ts.CalculateTax()

	OriginalAfterTax := [][]float64{} //, 0, len(ts.TaxRates)*len(ts.Prices))
    
	CodeTaxRate := ts.AfterTax

	for _, taxrate := range ts.TaxRates {
        tempValues := []float64{}
		for _, price := range ts.Prices {
            tempValues = append(tempValues, price * (1 + (taxrate/100)))
		}
		OriginalAfterTax = append(OriginalAfterTax, tempValues)
	}

	// if CodeTaxRate == OriginalAfterTax {
	// 	t.Errorf("Abs(-1) = %d; want 1", got)
	// }

	if !reflect.DeepEqual(CodeTaxRate, OriginalAfterTax) {
        t.Errorf("Invalid logic")
	}



}