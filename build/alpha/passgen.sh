#!/bin/bash

passwd=$(dd if=/dev/urandom count=1 2> /dev/null | uuencode -m - | sed -ne 2p | cut -c-12)

echo $passwd
#mosquitto_passwd -b /opt/tmp/mqtt_passwd test $passwd