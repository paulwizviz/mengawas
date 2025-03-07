package main

import (
	"log"
	"os"
)

func main() {
	// Get InfluxDB credentials from environment variables
	influxCred := os.Getenv("INFLUXDB_TOKEN")
	if influxCred == "" {
		log.Fatal("InfluxDB token is not specified in INFLUXDB_TOKEN")
	}

	influxBucket := os.Getenv("DOCKER_INFLUXDB_INIT_BUCKET")
	if influxBucket == "" {
		log.Fatal("InfluxDB token is not specified in DOCKER_INFLUXDB_INIT_BUCKET")
	}

	influxOrg := os.Getenv("DOCKER_INFLUXDB_INIT_ORG")
	if influxOrg == "" {
		log.Fatal("InfluxDB token is not specified in DOCKER_INFLUXDB_INIT_ORG")
	}
}
