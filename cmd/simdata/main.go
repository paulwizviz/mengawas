package main

import (
	"log"
	"mengawas/internal/iot/temperature"
	"os"
	"path/filepath"
	"time"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tempdatafile := filepath.Join(pwd, "testdata", "tempdata.csv")
	if err := temperature.SimulateDataToCSV(tempdatafile, 100, 20, 30, time.Now(), time.Now().Add(5*24*time.Hour), 6, 5, temperature.Celsius); err != nil {
		log.Fatal(err)
	}
}
