package sinks

import (
	appd "appdynamics"
	"log"
)

type ControllerClient struct {
	logger     *log.Logger
	regMetrics map[string]bool
}

func NewControllerClient(host, accessKey, account, app, tier, node string, port uint16, useSSL bool, logger *log.Logger) *ControllerClient {
	cfg := appd.Config{}

	cfg.AppName = app
	cfg.TierName = tier
	cfg.NodeName = node
	cfg.Controller.Host = host
	cfg.Controller.Port = port
	cfg.Controller.UseSSL = useSSL
	cfg.Controller.Account = account
	cfg.Controller.AccessKey = accessKey
	cfg.InitTimeoutMs = 1000
	appd.InitSDK(&cfg)
	logger.Println(&cfg.Controller)
	return &ControllerClient{logger: logger, regMetrics: make(map[string]bool)}
}

func (c *ControllerClient) PostBatch(events []interface{}) error {
	bt := appd.StartBT("PostBatch", "")
	for _, event := range events {
		if event != nil {
			dataPoint, ok := event.(*DataPoint)
			if !ok {
				continue
			}

			if dataPoint.Allowed {
				_, pres := c.regMetrics[dataPoint.Metric]
				if !pres {
					c.logger.Printf("Registering Metric: %v", dataPoint.Metric)
					appd.AddCustomMetric("", dataPoint.Metric, appd.APPD_TIMEROLLUP_TYPE_AVERAGE,
						appd.APPD_CLUSTERROLLUP_TYPE_INDIVIDUAL, appd.APPD_HOLEHANDLING_TYPE_REGULAR_COUNTER)
					c.regMetrics[dataPoint.Metric] = true
				}
				appd.ReportCustomMetric("", dataPoint.Metric, dataPoint.Value)
			}
		}
	}
	appd.EndBT(bt)

	return nil
}
