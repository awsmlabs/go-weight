package weight

import (
	"fmt"
	"strings"

	"go.awsmlabs.co.uk/rounding"
)

type Weight float64

const (
	Kilogram Weight = 1
	Newton          = Kilogram / 9.81
	Gram            = Kilogram / 1000
	Ounce           = Gram * 28.3495
	Pound           = Ounce * 16
	Stone           = Pound * 14
	Tonne           = Kilogram * 1000
	Carat           = Kilogram / 5000
)

func (w Weight) String() string {
	return fmt.Sprintf("%vkg", rounding.RoundNearest(float64(w), 2))
}

func (w Weight) Kilograms() float64 {
	return rounding.RoundNearest(float64(w), 2)
}

func (w Weight) Newtons() float64 {
	return rounding.RoundNearest(float64(w/Newton), 2)
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

func (w *Weight) MarshalJSON() ([]byte, error) {
	return []byte(w.String()), nil
}
