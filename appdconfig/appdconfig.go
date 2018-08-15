package appdconfig

import "github.com/kelseyhightower/envconfig"

type AppDConfig struct {
	Sink           string `default:"Controller" envconfig:"sink"`
	SamplingRate   uint16 `default:"2" envconfig:"sampling_rate"`
	ControllerHost string `envconfig:"controller_host"`
	ControllerPort uint16 `default:"8090" envconfig:"controller_port"`
	AccessKey      string `envconfig:"access_key"`
	Account        string `envconfig:"account"`
	SslEnabled     bool   `default:"false" envconfig:"ssl_enabled"`
	NozzleAppName  string `default:"appd-nozzle" envconfig:"nozzle_app"`
	NozzleTierName string `default:"appd-nozzle-tier" envconfig:"nozzle_tier"`
	NozzleNodeName string `default:"appd-nozzle-node" envconfig:"nozzle_node"`
}

func Parse() (*AppDConfig, error) {
	config := &AppDConfig{}
	err := envconfig.Process("appd", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
