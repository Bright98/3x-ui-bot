#!/bin/bash

$(docker stop $(docker ps -a -q --filter ancestor=3x-ui-bot --format="{{.ID}}"))
docker build -t 3x-ui-bot .
docker run -d 3x-ui-bot