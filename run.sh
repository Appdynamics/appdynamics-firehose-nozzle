#!/usr/bin/env bash

echo "Doing UAA Auth"

export NOZZLE_UAA_URL=<uaa address>   # cf curl /v2/info | jq .doppler_logging_endpoint
export NOZZLE_TRAFFIC_CONTROLLER_URL=<doppler address>  # cf curl /v2/info | jq .token_endpoint
export NOZZLE_USERNAME=<opentsb_firehose_username>
export NOZZLE_PASSWORD=<opentsb_firehose_password>
export NOZZLE_SINK=stdout # one of the following stdout|controller
export NOZZLE_FIREHOSE_SUBSCRIPTION_ID=appdynamics.firehose
export NOZZLE_SKIP_SSL=true
export NOZZLE_SELECTED_EVENTS=ValueMetric,CounterEvent
export APPD_SAMPLING_RATE: 2 
export APPD_CONTROLLER_HOST: <controller>
export APPD_ACCESS_KEY: <accesskey>
export APPD_CONTROLLER_PORT: <port>
export APPD_ACCOUNT: <account>


go run main.go
