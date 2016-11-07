#!/usr/bin/env bash

go run xke-app.go -clientIp=10.0.0.2:4444 -ourIp="${OUR_IP}":4242 -bartenderIp=10.0.0.2:4343 -playerId="${OUR_NAME}"