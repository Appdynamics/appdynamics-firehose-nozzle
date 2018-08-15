package main

import (
	"crypto/tls"
	"log"
	"os"
	"time"
	"errors"
	"strings"

	"github.com/cloudfoundry/noaa/consumer"

	"github.com/Appdynamics/firehose-utils/api"
	"github.com/Appdynamics/firehose-utils/config"
	"github.com/Appdynamics/firehose-utils/nozzle"
	"github.com/Appdynamics/firehose-utils/uaa"
	"github.com/Appdynamics/firehose-utils/writernozzle"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/sinks"

)

func main() {
	logger := log.New(os.Stdout, ">>> ", 0)

	conf, err := config.Parse()
	if err != nil {
		logger.Fatal("Unable to build config from environment", err)
	}

	var token, trafficControllerURL string
	if conf.APIURL != "" {
		logger.Printf("Fetching auth token via API: %v\n", conf.APIURL)

		fetcher, err := api.NewAPIClient(conf.APIURL, conf.Username, conf.Password, conf.SkipSSL)
		if err != nil {
			logger.Fatal("Unable to build API client", err)
		}
		token, err = fetcher.FetchAuthToken()
		if err != nil {
			logger.Fatal("Unable to fetch token via API", err)
		}

		trafficControllerURL = fetcher.FetchTrafficControllerURL()
		if trafficControllerURL == "" {
			logger.Fatal("trafficControllerURL from client was blank")
		}
	} else if conf.UAAURL != "" {
		logger.Printf("Fetching auth token via UAA: %v\n", conf.UAAURL)

		trafficControllerURL = conf.TrafficControllerURL
		if trafficControllerURL == "" {
			logger.Fatal(errors.New("NOZZLE_TRAFFIC_CONTROLLER_URL is required when authenticating via UAA"))
		}

		fetcher := uaa.NewUAATokenFetcher(conf.UAAURL, conf.Username, conf.Password, conf.SkipSSL)
		token, err = fetcher.FetchAuthToken()
		if err != nil {
			logger.Fatal("Unable to fetch token via UAA", err)
		}
	} else {
		logger.Fatal(errors.New("One of NOZZLE_API_URL or NOZZLE_UAA_URL are required"))
	}

	logger.Printf("Consuming firehose: %v\n", trafficControllerURL)
    noaaConsumer := consumer.New(trafficControllerURL, &tls.Config{
		InsecureSkipVerify: conf.SkipSSL,
	}, nil)
    eventsChan, errsChan := noaaConsumer.Firehose(conf.FirehoseSubscriptionID, token)
    
	sink := strings.ToLower(os.Getenv("NOZZLE_SINK"))
    
    var eventSerializer nozzle.EventSerializer
    var sinkWriter nozzle.Client
    switch sink {
	case sinks.Stdout:
            eventSerializer = writernozzle.NewWriterEventSerializer()
            sinkWriter = writernozzle.NewWriterClient(os.Stdout)
        case sinks.MachineAgent:
            logger.Fatal(errors.New("Not Implemented!"))
		case sinks.Controller:
			port := uint16(8090)
			host, accesskey, account := os.Getenv("APPD_CONTROLLER"), os.Getenv("APPD_ACCESSKEY"), os.Getenv("APPD_ACCOUNT")
			useSSL := false	
			sinkWriter = sinks.NewControllerClient(host, accesskey, account, port, useSSL)
            eventSerializer = sinks.NewControllerEventSerializer()
		default:
            logger.Fatal(errors.New("set NOZZLE_SINK environment variable to one of the following STDOUT|MACHINEAGENT|CONTROLLER|SPLUNK"))
    }

	  
    logger.Printf("Forwarding events to %s: %s", sink, conf.SelectedEvents)
	
    forwarder := nozzle.NewForwarder(sinkWriter, eventSerializer,
     conf.SelectedEvents, eventsChan, errsChan, logger)
    err = forwarder.Run(time.Second)
	if err != nil {
		logger.Fatal("Error forwarding", err)
	}
}
