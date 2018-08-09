
# Steps



## Setup GoPATH


```
mkdir workspace; cd workspace
export GOPATH=`pwd`

```

## Fetch the source

```
go get github.com/Appdynamics/appdynamics-firehose-nozzle
```


## Install Glide

```
curl https://glide.sh/get | sh
```

## Build

```
cd src/github.com/Appdynamics/appdynamics-firehose-nozzle
glide install
go install
```

## Run


```
Edit run.sh accordingly
source run.sh
```

## How to find Credentials


- NOZZLE_TRAFFIC_CONTROLLER_URL: wss://doppler.sys.pie-20.cfplatformeng.com:443 [https://docs.pivotal.io/pivotalcf/2-2/loggregator/architecture.html#components]

```
(master)$ cf curl /v2/info | jq .doppler_logging_endpoint
"wss://doppler.sys.pie-20.cfplatformeng.com:443"
```

- NOZZLE_UAA_URL: https://uaa.{domain from the above ie.. (sys.pie-20.cfplatformeng.com)}  

- NOZZLE_USERNAME/NOZZLE_PASSWORD: 
   * easy way: login to opsmanager and go to PAS tile -> credentials -> UAA -> Opentsdb Nozzle Credentials - Copy the username and password
   * (or) create a new account with permissions, https://github.com/cf-platform-eng/firehose-nozzle#option-2-uaa-client

   
