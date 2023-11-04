#!/bin/bash

COMMAND="$1"

export DEVICE_IMAGE_NAME=mengawas/devicemock:current
export DEVPROC_IMAGE_NAME=mengawas/devprocmock:current
export DEVOP_IMAGE_NAME=mengawas/devop:current

export MOSQUITTO_CONTAINER="mosquitto-server"
export MOCK_DEVICE_CONTAINER="mock-device"
export MOCK_DEVPROC_CONTAINER="mock-devproc"
export DEVOP_CONTAINER="devop"

export CLIENT_APP_NAME="mqttclient"

export NETWORK_NAME="mengawas_mqtt-network"

function build(){
    docker-compose -f ./build/ex1/builder.yaml build device
    docker-compose -f ./build/ex1/builder.yaml build devproc
    docker-compose -f ./build/ex1/builder.yaml build devop
}

function clean(){
    docker rmi -f ${DEVICE_IMAGE_NAME}
    docker rmi -f ${DEVPROC_IMAGE_NAME}
    docker rmi -f ${DEVOP_IMAGE_NAME}
    docker rmi -f $(docker images --filter "dangling=true" -q)
}

function run(){
    docker-compose -f ./deployment/ex1/docker-compose.yml up
}

function devop(){
    docker-compose -f ./deployment/ex1/devop.yml up passwd
}

function stop(){
    docker-compose -f ./deployment/ex1/docker-compose.yml down
}

message="$0 build | clean | devop | run  | stop"

if [ "$#" -ne 1 ]; then
    echo $message
    exit 1
fi

case $COMMAND in
    "build")
        build
        ;;
    "clean")
        clean
        ;;
    "devop")
        devop
        ;;
    "run")
        run
        ;;
    "stop")
        stop
        ;;
    *)
        echo $message
        ;;
esac