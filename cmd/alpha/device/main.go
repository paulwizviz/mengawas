package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"time"

	"mengawas/internal/nmodel"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fxamacker/cbor/v2"
)

func connect(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	log.Printf("===> clientID: %v password:%v", clientId, password)
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	// Get a publishing client
	client := connect("pub", uri)
	// Create a timer that ticks every second
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		d := nmodel.Data{
			ID:      nmodel.CreateID(),
			Content: []byte(fmt.Sprintf("Time: %v", t)),
		}
		// Serialize model
		cborD, _ := cbor.Marshal(d)
		// Publish
		client.Publish(topic, 0, false, cborD)
	}
}
