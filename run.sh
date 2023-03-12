#!/bin/bash

on_exit(){
  pkill weaver
}

CODE_PATH=./src

COMMAND="go run $CODE_PATH"
MODE="single"

if [ "$1" == "multi" ]; then
  MODE="multi"
  COMMAND="weaver "$MODE" deploy weaver.toml"
fi

trap 'on_exit' EXIT

go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
go mod tidy
weaver generate $CODE_PATH
go build -o ./go-service-weaver-tutorial $CODE_PATH
$COMMAND &
sleep 5
ab -c 50 -n 5000 localhost:12345/hello?name=Walter%20White &
P1=$!
weaver $MODE dashboard &
wait $P1
weaver $MODE status
read -p "Press a button to exit..." -r nothing
