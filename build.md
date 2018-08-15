
# Steps to build the firehose nozzle and run



## Make sure to install GoLang 1.8+

https://golang.org/doc/install


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

