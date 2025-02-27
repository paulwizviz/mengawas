#!/bin/bash

if [ "$(basename $(realpath .))" != "mengawas" ]; then
    echo "You are outside the scope of the project"
    exit 0
fi

COMMAND="$1"
SUBCOMMAND="$2"

export MOSQUITTO_IMAGE=eclipse-mosquitto:2.0.14
export MOSQUITTO_CONTAINER="mosquitto-server"

export INFLUX_IMAGE=influxdb:2.6

export PGADMIN_IMAGE=dpage/pgadmin4:8.9
export PLSQL_IMAGE=postgres:16.2-alpine3.19

export NETWORK=mengaws_vineyard

function start_network(){
    docker compose -f ./deployment/vineyard/mqtt.yaml \
                   -f ./deployment/vineyard/influx.yaml \
                   -f ./deployment/vineyard/psql.yaml up
}

function stop_network(){
    docker compose -f ./deployment/vineyard/mqtt.yaml \
                   -f ./deployment/vineyard/influx.yaml \
                   -f ./deployment/vineyard/psql.yaml down
}

function clean(){
    stop_network
    docker rmi -f ${MOSQUITTO_IMAGE}
    docker rmi -f ${INFLUX_IMAGE}
    docker rmi -f ${PLSQL_IMAGE}
    docker rmi -f ${PGADMIN_IMAGE}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

case $COMMAND in
    "clean")
        clean
        ;;
    "start")
        start_network
        ;;
    "stop")
        stop_network
        ;;
    *)
        echo "Usage: $0 [start]"
        ;;
esac