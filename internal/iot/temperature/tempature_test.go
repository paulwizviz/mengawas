package temperature

import (
	"fmt"
	"mengawas/internal/iot"
	"testing"
	"time"

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

func TestNewMeasureTimeStamp(t *testing.T) {
	testcases := []struct {
		name    string
		input   any
		want    Measurement
		wantErr error
	}{
		{
			name: "NewMeasure",
			input: func() time.Time {
				return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
		},
		{
			name: "NewMeasureS",
			input: func() string {
				return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.RFC3339)
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
			wantErr: nil,
		},
		{
			name: "NewMeasureS",
			input: func() string {
				return "2025-01-01T00:00:00Z" // UTC format
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
			wantErr: nil,
		},
		{
			name: "NewMeasureS",
			input: func() string {
				return "2025-01-01T00:00:00" // Non-UTC format
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Time{}
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
			wantErr: iot.ErrNewMeasure,
		},
		{
			name: "NewMeasureS",
			input: func() string {
				return "2025-01-01T00:00:00+01:00" // Local timestamp format
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Time{}
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
			wantErr: iot.ErrNotUTC,
		},
		{
			name: "NewMeasureS",
			input: func() string {
				return "abcd" // Invalid time format
			}(),
			want: Measurement{
				location: "location",
				deviceID: "deviceID",
				timeStamp: func() time.Time {
					return time.Time{}
				}(),
				unit: Unit{
					value: 70,
					unit:  Celsius,
				},
			},
			wantErr: iot.ErrNewMeasure,
		},
	}

	for i, tc := range testcases {
		t.Run(fmt.Sprintf("%d-%v", i, tc.name), func(t *testing.T) {
			location := tc.want.location
			deviceID := tc.want.deviceID
			unit := NewCelsius(tc.want.unit.value)
			switch ts := tc.input.(type) {
			case time.Time:
				got := NewMeasure(location, deviceID, ts, unit)
				assert.Equal(t, tc.want.Timestamp(), got.Timestamp(), "Not the same time stamp")
				assert.Equal(t, time.UTC, got.Timestamp().Location(), "Not UTC")
			case string:
				got, gotErr := NewMeasureS(location, deviceID, ts, NewCelsius(tc.want.unit.value))
				if assert.ErrorIs(t, gotErr, tc.wantErr, "error asserted") {
					assert.Equal(t, tc.want.Timestamp(), got.Timestamp(), "Same time")
				}
			}
		})
	}
}
