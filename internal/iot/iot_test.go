package iot

import (
	"fmt"
	"mengawas/internal/iot/temperature"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ts time.Time = time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

type testable interface {
	Test(t *testing.T)
}

type newMeasureTest[U UnitOfMeasure] struct {
	input U
	want  Measurement[U]
}

func (n newMeasureTest[U]) Test(t *testing.T) {
	switch any(n.input).(type) {
	case temperature.Unit:
		got := NewMeasurement("location", "deviceID", "measure type", ts, n.input)
		assert.Equal(t, n.want, got, fmt.Sprintf("Want: %v Got: %v", n.want, got))
	}
}

func TestNewMeasure(t *testing.T) {
	testcases := []testable{
		newMeasureTest[temperature.Unit]{
			input: temperature.NewCelsius(70),
			want: Measurement[temperature.Unit]{
				location:    "location",
				deviceID:    "deviceID",
				measureType: "measure type",
				timeStamp:   ts,
				unit:        temperature.NewCelsius(70),
			},
		},
	}
	for i, tc := range testcases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tc.Test(t)
		})
	}
}
