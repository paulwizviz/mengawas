
ARG UBUNTU_VER

FROM ubuntu:${UBUNTU_VER}

COPY ./build/ex1/passgen.sh /opt/passgen.sh

RUN apt-get update && \
    apt-get install -y \
            sharutils \
            mosquitto \
            mosquitto-clients 
