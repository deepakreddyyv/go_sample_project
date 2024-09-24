package tax

import (
	"example.com/price_calculator/conversions"
	"example.com/price_calculator/iomanager"
)


type TaxStore struct {
    IoManager iomanager.IoManager `json:"-"`
	TaxRates []float64 `json:"tax_rates"`
	Prices []float64   `json:"prices"`
	AfterTax [][]float64 `json:"after_tax"`
}

func (t *TaxStore) TransformValues() error {
	prices := []float64{}
    
	contents, err := t.readValues() 
	if err != nil {
		return err
	}

	for _, val := range contents{
		for _, val2 := range val {
			convInit := conversions.StringData{
				Data: val2,
				Base: 10,
				BitSize: 64,
			}// strconv.ParseFloat(val2, 64)
			value, err := convInit.ToFloat()
            if err != nil {
				return err
			}
            prices = append(prices, value)
		}
	}

	t.Prices = prices
	return nil
}

func (t *TaxStore) CalculateTax() error {
	var afterTax [][]float64 = [][]float64{}

	for _, tax := range t.TaxRates {
		taxCompute := []float64{}
		for _, prices := range t.Prices {
			res := prices * (1 + (tax / 100))
            taxCompute = append(taxCompute, res)
		}
		afterTax = append(afterTax, taxCompute)
	}
	t.AfterTax = afterTax

	return t.writeValues()
}

func (t *TaxStore) readValues() ([][]string, error){
	values, err := t.IoManager.Read()

	if err != nil {
		return [][]string{}, err
	}

	return values, nil 
}

func (t *TaxStore) writeValues() error {
	err := t.IoManager.Write(t)

	if err != nil {
		return err 
	}

	return nil
}



