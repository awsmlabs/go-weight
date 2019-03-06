package weight

import (
	"fmt"
	rounding "github.com/mealfitx/rounding"
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

func (d Weight) String() string {
	return fmt.Sprintf("%vkg", rounding.RoundNearest(float64(d), 2))
}

func (d Weight) Kilograms() float64 {
	return rounding.RoundNearest(float64(d), 2)
}

func (d Weight) Newtons() float64 {
	return rounding.RoundNearest(float64(d / Newton), 2)
}

func (d Weight) Ounces() float64 {
	return rounding.RoundNearest(float64(d / Ounce), 2)
}

func (d Weight) Pounds() float64 {
	return rounding.RoundNearest(float64(d / Pound), 2)
}

func (d Weight) Stones() float64 {
	return rounding.RoundNearest(float64(d / Stone), 2)
}

func (d Weight) Tonnes() float64 {
	return rounding.RoundNearest(float64(d / Tonne), 2)
}

func (d Weight) Carats() float64 {
	return rounding.RoundNearest(float64(d / Carat), 2)
}

