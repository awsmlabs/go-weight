package weight

import (
	"testing"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		in       Weight
		expected string
	}{
		{5 * Kilogram, "5kg"},
		{2.2 * Pound, "1kg"},
		{165 * Pound, "74.84kg"},
		{1000 * Gram, "1kg"},
		{9.8 * Newton, "1kg"},
		{5000 * Carat, "1kg"},
		{1 * Tonne, "1000kg"},
	}

	for _, tc := range testCases {
		if out := tc.in.String(); out != tc.expected {
			t.Errorf("Wrong string, got %v expected %v", out, tc.expected)
		}
	}
}

func TestParseWeightToKilograms(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"10kg", 10},
		{"-5kg", -5},
		{"2.2lb", 1},
		{"165lb", 74.84},
		{"12st", 76.2},
		{"12st5lb", 78.47},
		{"2kg1000g", 3},
		{"2kg500g", 2.5},
		{"160lb0oz", 72.57},
		{"160lb8oz", 72.8},
		{"161lb", 73.03},
		{"8oz", 0.23},
		{"150N", 15.29},
		{"5000c", 1},
		{"1tn", 1000},

	}

	for _, tc := range testCases {
		calculated, err := ParseWeight(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if calculated.Kilograms() != tc.expected {
			t.Errorf("Wrong distance, got %v expected %v", calculated.Kilograms(), tc.expected)
		}
	}
}

func TestParseWeightToPounds(t *testing.T) {
	testCases := []struct {
		in       string
		expected float64
	}{
		{"2.2lb", 2.2},
		{"1st", 14},
		{"1kg", 2.2},
		{"1tn", 2204.62},
		{"1st7lb", 21},
		{"1st7lb8oz", 21.5},
	}

	for _, tc := range testCases {
		calculated, err := ParseWeight(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}

		if calculated.Pounds() != tc.expected {
			t.Errorf("Wrong distance, got %v expected %v", calculated.Pounds(), tc.expected)
		}
	}
}
