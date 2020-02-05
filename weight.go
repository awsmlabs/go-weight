package weight

import (
	"strconv"
	"strings"

	"go.awsmlabs.co.uk/rounding"
)

type Weight float64

// weights
const (
	Microgram Weight = 1
	Milligram        = Microgram * 1000
	Carat            = Milligram * 200
	Gram             = Milligram * 1000
	Kilogram         = Gram * 1000
	Tonne            = Kilogram * 1000
	Ounce            = Gram * 28.3495
	Pound            = Ounce * 16
	Stone            = Pound * 14
	Newton           = Kilogram / 9.81
)

func (w Weight) String() string {
	switch {
	case w < Gram:
		return strconv.FormatFloat(rounding.RoundNearest(w.Milligrams(), 2), 'f', -1, 64) + "mg"
	case w < Kilogram:
		return strconv.FormatFloat(rounding.RoundNearest(w.Grams(), 2), 'f', -1, 64) + "g"
	case w < Tonne:
		return strconv.FormatFloat(rounding.RoundNearest(w.Kilograms(), 2), 'f', -1, 64) + "kg"
	default:
		return strconv.FormatFloat(rounding.RoundNearest(w.Tonnes(), 2), 'f', -1, 64) + "tn"
	}
}

func (w Weight) Kilograms() float64 {
	return rounding.RoundNearest(float64(w/Kilogram), 2)
}

func (w Weight) Newtons() float64 {
	return rounding.RoundNearest(float64(w/Newton), 2)
}

func (w Weight) Grams() float64 {
	return rounding.RoundNearest(float64(w/Gram), 2)
}

func (w Weight) Milligrams() float64 {
	return rounding.RoundNearest(float64(w/Milligram), 2)
}

func (w Weight) Ounces() float64 {
	return rounding.RoundNearest(float64(w/Ounce), 2)
}

func (w Weight) Pounds() float64 {
	return rounding.RoundNearest(float64(w/Pound), 2)
}

func (w Weight) Stones() float64 {
	return rounding.RoundNearest(float64(w/Stone), 2)
}

func (w Weight) Tonnes() float64 {
	return rounding.RoundNearest(float64(w/Tonne), 2)
}

func (w Weight) Carats() float64 {
	return rounding.RoundNearest(float64(w/Carat), 2)
}

func (w *Weight) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}

	*w, err = ParseWeight(s)
	return
}

func (w Weight) MarshalJSON() ([]byte, error) {
	return []byte(`"` + w.String() + `"`), nil
}
