package iot

import (
	"errors"
	"mengawas/internal/iot/acidity"
	"mengawas/internal/iot/temperature"
	"time"

	"github.com/fxamacker/cbor/v2"
)

var (
	ErrNewMeasure         = errors.New("unable to instantiate new measurement")
	ErrNotUTC             = errors.New("not valud UTC")
	ErrInvalidMeasureCBOR = errors.New("invalid measurement cbor")
)

type UnitOfMeasure interface {
	temperature.Unit | acidity.Unit
}

type Measurement[U UnitOfMeasure] struct {
	location     string
	deviceID     string
	measureGroup string
	timeStamp    time.Time
	unit         U
}

func (m Measurement[U]) Location() string {
	return m.location
}

func (m Measurement[U]) DeviceID() string {
	return m.deviceID
}

func (m Measurement[U]) Timestamp() time.Time {
	return m.timeStamp
}

func (m Measurement[U]) Type() string {
	return m.measureGroup
}

func (m Measurement[U]) Unit() U {
	return m.unit
}

type aux[U UnitOfMeasure] struct {
	Location  string    `cbor:"location"`
	DeviceID  string    `cbor:"device_id"`
	Type      string    `cbor:"type"`
	TimeStamp time.Time `cbor:"timestamp"`
	Unit      U         `cbor:"unit"`
}

func (m Measurement[U]) MarshalCBOR() ([]byte, error) {
	aux := aux[U]{
		Location:  m.location,
		DeviceID:  m.deviceID,
		Type:      m.measureGroup,
		TimeStamp: m.timeStamp,
		Unit:      m.unit,
	}
	return cbor.Marshal(aux)
}

func (m *Measurement[U]) UnmarshalCBOR(data []byte) error {
	var aux aux[U]
	err := cbor.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	m.location = aux.Location
	m.deviceID = aux.DeviceID
	m.measureGroup = aux.Type
	m.timeStamp = aux.TimeStamp
	m.unit = aux.Unit
	return nil
}

func NewMeasurement[U UnitOfMeasure](location string, deviceID string, measureGroup string, ts time.Time, u U) Measurement[U] {
	return Measurement[U]{
		location:     location,
		deviceID:     deviceID,
		measureGroup: measureGroup,
		timeStamp:    ts,
		unit:         u,
	}
}
