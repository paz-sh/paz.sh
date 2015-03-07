#!/bin/bash

PLATFORM=`uname -s | tr '[:upper:]' '[:lower:]'`
ARCH=`uname -m`

if [ "$ARCH" = "x86_64" ]; then
    ARCH="amd64"
else
    ARCH="386"
fi

PAZ_COMMAND="./bin/paz_${PLATFORM}-${ARCH}"
if ! which "${PAZ_COMMAND}" &>/dev/null; then
  PAZ_COMMAND="./bin/paz"
fi

function paz() {
    "${PAZ_COMMAND}" $@
}
