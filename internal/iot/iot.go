package iot

import "errors"

var (
	ErrNewMeasure         = errors.New("unable to instantiate new measurement")
	ErrNotUTC             = errors.New("not valud UTC")
	ErrInvalidMeasureCBOR = errors.New("invalid measurement cbor")
)

var (
	TypeTemperature = "TemperateMeasurement"
	TypeWind        = "WindMeasurement"
	TypeSoil        = "SoilMeasurement"
)
