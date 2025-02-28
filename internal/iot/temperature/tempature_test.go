package temperature

import (
	"fmt"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	testcases := []struct {
		name  string
		value float32
		want  Unit
	}{
		{
			name:  "celsius",
			value: 70.5,
			want: Unit{
				value: 70.5,
				unit:  Celsius,
			},
		},
		{
			name:  "fahrenheit",
			value: 70.5,
			want: Unit{
				value: 70.5,
				unit:  Fahrenheit,
			},
		},
		{
			name:  "kelvin",
			value: 70.5,
			want: Unit{
				value: 70.5,
				unit:  Kelvin,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "celsius" {
				got := NewCelsius(tc.value)
				assert.Equal(t, tc.want, got, "expected celelsius")
			} else if tc.name == "fahrenheit" {
				got := NewFahrenheit(tc.value)
				assert.Equal(t, tc.want, got, "expected fahrenheit")
			} else {
				got := NewKelvin(tc.value)
				assert.Equal(t, tc.want, got, "expected fahrenheit")
			}
		})
	}
}

func TestSerialiseUnit(t *testing.T) {
	testcases := []struct {
		name string
		want Unit
	}{
		{
			name: "celsius",
			want: Unit{
				value: 70,
				unit:  Celsius,
			},
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if tc.name == Celsius {
				c := NewCelsius(70)
				b, _ := cbor.Marshal(c)
				var got Unit
				cbor.Unmarshal(b, &got)
				assert.Equal(t, tc.want, got, fmt.Sprintf("Want: %v Got: %v", tc.want, got))
			}
		})
	}
}
