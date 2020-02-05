package weight

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToString(t *testing.T) {
	testCases := []struct {
		in       Weight
		expected string
	}{
		{5 * Kilogram, "5kg"},
		{2.21 * Pound, "1kg"},
		{165 * Pound, "74.84kg"},
		{1000 * Gram, "1kg"},
		{9.81 * Newton, "1kg"},
		{5000 * Carat, "1kg"},
		{1 * Tonne, "1tn"},
	}

	for _, tc := range testCases {
		if out := tc.in.String(); out != tc.expected {
			t.Errorf("Wrong string, got %v expected %v", out, tc.expected)
		}
	}
}

func TestJSONUnmarshal(t *testing.T) {
	type testStruct struct {
		Field1 Weight `json:"field_1"`
		Field2 Weight `json:"field_2"`
		Field3 Weight `json:"field_3"`
	}

	testData := &testStruct{}
	b := []byte(`{"field_1":"10g","field_2":"100kg"}`)

	if assert.Nil(t, json.Unmarshal(b, testData), "failed to unmarshal") {
		assert.Equal(t, 10*Gram, testData.Field1)
		assert.Equal(t, 100*Kilogram, testData.Field2)
		assert.Equal(t, 0*Gram, testData.Field3)
	}
}

func TestJSONMarshal(t *testing.T) {
	testData := struct {
		Field1 Weight `json:"field_1,omitempty"`
		Field2 Weight `json:"field_2,omitempty"`
		Field3 Weight `json:"field_3,omitempty"`
	}{
		Field1: 10 * Gram,
	}

	b, err := json.Marshal(testData)
	assert.Nil(t, err)
	assert.Equal(t, `{"field_1":"10g"}`, string(b))
}
