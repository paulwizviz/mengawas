package temperature

import "github.com/fxamacker/cbor/v2"

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

type aux struct {
	Value float32 `cbor:"value"`
	Unit  string  `cbor:"unit"`
}

func (u Unit) MarshalCBOR() ([]byte, error) {
	aux := aux{
		Value: u.value,
		Unit:  u.unit,
	}
	return cbor.Marshal(aux)
}

func (u *Unit) UnmarshalCBOR(data []byte) error {
	var aux aux
	err := cbor.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	u.value = aux.Value
	u.unit = aux.Unit
	return nil
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
