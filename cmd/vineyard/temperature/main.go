package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"mengawas/internal/iot"
	"mengawas/internal/iot/temperature"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fxamacker/cbor/v2"
)

func readCSV(r io.Reader) (<-chan iot.Measurement[temperature.Unit], error) {
	ch := make(chan iot.Measurement[temperature.Unit], 1)
	csvr := csv.NewReader(r)
	_, err := csvr.Read()
	if err != nil {
		return nil, err
	}
	go func() {
		defer close(ch)
	loop:
		for {
			data, err := csvr.Read()
			if err != nil {
				if errors.Is(err, io.EOF) {
					break loop
				}
				log.Println(err)
				continue loop
			}
			ts, err := time.Parse(time.RFC3339, data[0])
			if err != nil {
				log.Println("parse ", data[0], err)
				continue loop
			}
			tempVal := data[1]
			value, err := strconv.ParseFloat(tempVal, 32)
			if err != nil {
				log.Println(err)
				continue loop
			}
			m := iot.NewMeasurement("location", "temp-device", "temperature", ts, temperature.NewCelsius(float32(value)))
			ch <- m
		}
	}()
	return ch, nil
}

func serialiseTemparature(temps <-chan iot.Measurement[temperature.Unit]) chan []byte {
	out := make(chan []byte, 1)
	go func() {
		for temp := range temps {
			b, err := cbor.Marshal(temp)
			if err != nil {
				log.Println(err)
				continue
			}
			out <- b
		}
		close(out)
	}()
	return out
}

func publishToBroker(ins ...chan []byte) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://localhost:1883")
	opts.SetClientID("clientID")
	topic := "test/hello"
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	for _, in := range ins {
		for payload := range in {
			if token := client.Publish(topic, 0, false, payload); token.Wait() && token.Error() != nil {
				log.Fatal(token.Error())
			}
			log.Printf("Publish to topic: %s %v", topic, payload)
		}
	}
}

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	tempdatafile := filepath.Join(pwd, "testdata", "tempdata.csv")

	f, err := os.Open(tempdatafile)
	if err != nil {
		log.Fatal(err)
	}

	in, err := readCSV(f)
	if err != nil {
		log.Fatal(err)
	}

	r1 := serialiseTemparature(in)
	r2 := serialiseTemparature(in)
	r3 := serialiseTemparature(in)

	publishToBroker(r1, r2, r3)

	go func() {
		opts := mqtt.NewClientOptions()
		opts.AddBroker("tcp://localhost:1883")
		opts.SetClientID("clientID")
		topic := "test/hello"
		client := mqtt.NewClient(opts)
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
		if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Println(msg)
		}); token.Wait() && token.Error() != nil {
			log.Fatal(token.Error())
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
