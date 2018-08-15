package sinks

import appd "appdynamics"

type ControllerClient struct{}

func NewControllerClient(host, accessKey, account string, port uint16, useSSL bool) *ControllerClient {
	cfg := appd.Config{}

	cfg.AppName = "appd-nozzle"
	cfg.TierName = "nozzle-tier"
	cfg.NodeName = "nozzle-node"
	cfg.Controller.Host = host
	cfg.Controller.Port = port
	cfg.Controller.UseSSL = useSSL
	cfg.Controller.Account = account
	cfg.Controller.AccessKey = accessKey
	cfg.InitTimeoutMs = 1000
	appd.InitSDK(&cfg)

	return &ControllerClient{}
}

func (c *ControllerClient) PostBatch(events []interface{}) error {
	bt := appd.StartBT("my bt", "")
	for _, event := range events {
		if event != nil {
			dataPoint, ok := event.(*DataPoint)
			if !ok {
				continue
			}
			if dataPoint.Source == "gorouter" || dataPoint.Source == "uaa" {
				appd.AddCustomMetric("", dataPoint.Metric, appd.APPD_TIMEROLLUP_TYPE_AVERAGE,
					appd.APPD_CLUSTERROLLUP_TYPE_INDIVIDUAL, appd.APPD_HOLEHANDLING_TYPE_REGULAR_COUNTER)

				appd.ReportCustomMetrics("", dataPoint.Metric, dataPoint.Value)
			}

		}
	}
	appd.EndBT(bt)

	return nil
}
