package appdconfig

import "github.com/kelseyhightower/envconfig"

type AppDConfig struct {
	Sink           string `default:"stdout" envconfig:"sink"`
	SamplingRate   uint16 `default:"2" envconfig:"sampling_rate"`
	ControllerHost string `envconfig:"controller_host"`
	ControllerPort uint16 `default: "8090" envconfig:"controller_port"`
	AccessKey      string `envconfig:"access_key"`
	Account        string `envconfig:"account"`
	SslEnabled     bool   `default:"false" envconfig:"ssl_enabled"`
}

func Parse() (*AppDConfig, error) {
	config := &AppDConfig{}
	err := envconfig.Process("appd", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
