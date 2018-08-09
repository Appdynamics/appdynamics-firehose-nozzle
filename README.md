
Steps
-----

```
mkdir workspace; cd workspace
export GOPATH=`pwd`

```

```
go get github.com/Appdynamics/appdynamics-firehose-nozzle
```

```
curl https://glide.sh/get | sh
```

```
cd src/github.com/Appdynamics/appdynamics-firehose-nozzle
glide install
go install
```

```
go run main.go
```
