package conversions

import "strconv"

type StringConv interface {
	ToInt() error
	ToFloat() error
}

type StringData struct {
	Data string
	Base uint8 
	BitSize int
}

func (s *StringData) ToFloat() (float64, error){
	value, err := strconv.ParseFloat(s.Data, s.BitSize)

	if err != nil {
		return 0.0, err 
	}

	return value, nil
}