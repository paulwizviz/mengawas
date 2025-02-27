package temperature

import (
	"fmt"
	"mengawas/internal/iot"
	"time"

	"github.com/fxamacker/cbor/v2"
)

var (
	Celsius    = "C"
	Fahrenheit = "F"
	Kelvin     = "K"
)

type Unit struct {
	value float32
	unit  string
}

func (u Unit) Value() float32 {
	return u.value
}

func (u Unit) Unit() string {
	return u.unit
}

func NewCelsius(value float32) Unit {
	return Unit{
		value: value,
		unit:  Celsius,
	}
}

func NewFahrenheit(value float32) Unit {
	return Unit{
		value: value,
		unit:  Fahrenheit,
	}
}

func NewKelvin(value float32) Unit {
	return Unit{
		value: value,
		unit:  Kelvin,
	}
}

type Measurement struct {
	location    string
	deviceID    string
	measureType string
	timeStamp   time.Time
	unit        Unit
}

func (m Measurement) Location() string {
	return m.location
}

func (m Measurement) DeviceID() string {
	return m.deviceID
}

func (m Measurement) Timestamp() time.Time {
	return m.timeStamp
}

func (m Measurement) Type() string {
	return m.measureType
}

func (m Measurement) Unit() Unit {
	return m.unit
}

type aux struct {
	Location  string  `cbor:"location"`
	DeviceID  string  `cbor:"deviceid"`
	Type      string  `cbor:"type"`
	TimeStamp int64   `cbor:"timestamp"`
	TempValue float32 `cbor:"tempvalue"`
	TempUnit  string  `cbor:"tempunit"`
}

func (m Measurement) MarshalCBOR() ([]byte, error) {
	tm := m.timeStamp.Unix()
	aux := aux{
		Location:  m.location,
		DeviceID:  m.deviceID,
		Type:      m.measureType,
		TimeStamp: tm,
		TempValue: m.unit.value,
		TempUnit:  m.unit.unit,
	}
	return cbor.Marshal(aux)
}

func (m *Measurement) UnmarshalCBOR(data []byte) error {
	var aux aux
	err := cbor.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	tm := time.Unix(aux.TimeStamp, 0)

	m.location = aux.Location
	m.deviceID = aux.DeviceID
	m.measureType = aux.Type
	m.timeStamp = tm
	m.unit = Unit{
		value: aux.TempValue,
		unit:  aux.TempUnit,
	}
	return nil
}

func NewMeasure(location string, deviceID string, timeStamp time.Time, unit Unit) Measurement {
	return Measurement{
		location:    location,
		deviceID:    deviceID,
		measureType: iot.TypeTemperature,
		timeStamp:   timeStamp,
		unit:        unit,
	}
}

func NewMeasureS(location string, deviceID string, utcTimeStamp string, unit Unit) (Measurement, error) {
	tm, err := time.Parse(time.RFC3339, utcTimeStamp)
	if err != nil {
		return Measurement{}, fmt.Errorf("%w-%v", iot.ErrNewMeasure, err)
	}
	if tm.Location() != time.UTC {
		return Measurement{}, fmt.Errorf("%w-%v", iot.ErrNotUTC, err)
	}
	return Measurement{
		location:  location,
		deviceID:  deviceID,
		timeStamp: tm,
		unit:      unit,
	}, err
}
