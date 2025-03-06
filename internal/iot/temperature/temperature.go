package temperature

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
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

func SimulateDataToCSV(filename string, numRecords int, minTemp, maxTemp float64, startTime, endTime time.Time, spikeFrequency float64, spikeAmplitude float64, tempUnit string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"timestamp", "temperature", "unit"}
	if err := writer.Write(header); err != nil {
		return err
	}

	duration := endTime.Sub(startTime)
	timeStep := duration / time.Duration(numRecords)
	currentTime := startTime

	for i := 0; i < numRecords; i++ {
		temperature := minTemp + rand.Float64()*(maxTemp-minTemp)

		if rand.Float64() < spikeFrequency {
			spike := spikeAmplitude * float64(rand.Intn(2)*2-1) // Randomly choose + or -
			temperature += spike
		}

		if temperature < minTemp {
			temperature = minTemp
		}
		if temperature > maxTemp {
			temperature = maxTemp
		}

		record := []string{currentTime.Format(time.RFC3339), fmt.Sprintf("%.2f", temperature), tempUnit}
		if err := writer.Write(record); err != nil {
			return err
		}
		currentTime = currentTime.Add(timeStep)
	}
	return nil
}
