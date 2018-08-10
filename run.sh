#!/usr/bin/env bash

auth_method="uaa"

if [[ "${auth_method}" == "api" ]]; then ## ignored
    echo "Doing API auth" ## ignored
    export NOZZLE_API_URL=https://api.bosh-lite.com ## ignored
    export NOZZLE_USERNAME=admin ## ignored
    export NOZZLE_PASSWORD=admin ## ignored

else ## CHANGE THE FOLLOWING 
    echo "Doing UAA Auth"
    export NOZZLE_UAA_URL=https://uaa.<domain>
    export NOZZLE_TRAFFIC_CONTROLLER_URL=<doppler address>
    export NOZZLE_USERNAME=<opentsb_firehose_username>
    export NOZZLE_PASSWORD=<opentsb_firehose_password>
fi

export NOZZLE_SINK=STDOUT # one of the following STDOUT|MACHINEAGENT|CONTROLLER|SPLUNK etc...
export NOZZLE_FIREHOSE_SUBSCRIPTION_ID=appdynamics.firehose
export NOZZLE_SKIP_SSL=true
export NOZZLE_SELECTED_EVENTS=ValueMetric,CounterEvent

go run main.go
