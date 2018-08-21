package main

import (
	"crypto/tls"
	"errors"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/appdconfig"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/config"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/nozzle"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/sinks"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/uaa"
	"github.com/Appdynamics/appdynamics-firehose-nozzle/writernozzle"
	"log"
	"os"
	"strings"
	"time"

	"github.com/cloudfoundry/noaa/consumer"
)

func main() {
	logger := log.New(os.Stdout, "[APPD_NOZZLE]>>>",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)

	conf, err := config.Parse()
	if err != nil {
		logger.Fatal("Unable to build Nozzle config from environment", err)
	}

	appdConf, err := appdconfig.Parse()
	if err != nil {
		logger.Fatal("Unable to build Appdynamics config from environment", err)
	}
	logger.Printf("Using Appdynamics Configuration %v", appdConf)

	var token, trafficControllerURL string

	if conf.UAAURL != "" {
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
		logger.Fatal(errors.New("NOZZLE_UAA_URL is required"))
	}

	logger.Printf("Consuming firehose: %v\n", trafficControllerURL)

	noaaConsumer := consumer.New(trafficControllerURL, &tls.Config{InsecureSkipVerify: conf.SkipSSL}, nil)
	eventsChan, errsChan := noaaConsumer.Firehose(conf.FirehoseSubscriptionID, token)

	var eventSerializer nozzle.EventSerializer
	var sinkWriter nozzle.Client

	switch strings.ToLower(appdConf.Sink) {
	case sinks.Stdout:
		eventSerializer = writernozzle.NewWriterEventSerializer()
		sinkWriter = writernozzle.NewWriterClient(os.Stdout)
	case sinks.Controller:
		sinkWriter = sinks.NewControllerClient(appdConf.ControllerHost,
			appdConf.AccessKey,
			appdConf.Account,
			appdConf.NozzleAppName,
			appdConf.NozzleTierName,
			appdConf.NozzleNodeName,
			appdConf.ControllerPort,
			appdConf.SslEnabled,
			logger)
		eventSerializer = sinks.NewControllerEventSerializer(appdConf.NozzleTierName)
	default:
		logger.Fatal(errors.New("set APPD_SINK environment variable to one of the following stdout|controller"))
	}

	logger.Printf("Forwarding events to %s: %s", appdConf.Sink, conf.SelectedEvents)

	flush_time := appdConf.SamplingRate
	forwarder := nozzle.NewForwarder(sinkWriter, eventSerializer,
		conf.SelectedEvents, eventsChan, errsChan, logger)

	err = forwarder.Run(time.Duration(flush_time) * time.Second)
	if err != nil {
		logger.Fatal("Error forwarding", err)
	}
}
