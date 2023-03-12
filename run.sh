#!/bin/bash

on_exit(){
  pkill weaver
}

COMMAND="go run ."
MODE="single"

if [ "$1" == "multi" ]; then
  MODE="multi"
  COMMAND="weaver "$MODE" deploy weaver.toml"
fi

trap 'on_exit' EXIT

go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
go mod tidy
weaver generate
go build
$COMMAND &
sleep 5
ab -c 50 -n 5000000 localhost:12345/hello?name=Walter%20White &
P1=$!
weaver $MODE dashboard &
wait $P1
