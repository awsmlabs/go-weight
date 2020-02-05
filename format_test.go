package weight

import (
	"testing"
)

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
