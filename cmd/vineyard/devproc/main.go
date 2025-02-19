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
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri *url.URL, topic string, c chan string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(_ mqtt.Client, msg mqtt.Message) {
		p := msg.Payload()
		// Unmarshal payload
		var d nmodel.Data
		err := cbor.Unmarshal(p, &d)
		if err != nil {
			log.Println(err)
		}
		c <- fmt.Sprintf("* [%s] +----+ %v\n", msg.Topic(), string(d.Content))
	})
}

func main() {

	urlPtr := flag.String("url", "tcp://mosquitto-server:1883/", "location of broker")
	topicPtr := flag.String("topic", "test/hello", "topic name")
	flag.Parse()

	s := *urlPtr
	topic := *topicPtr

	uri, _ := url.Parse(s)

	// Listen to a topic
	c := make(chan string)
	go listen(uri, topic, c)

	for {
		fmt.Println(<-c)
	}

}
