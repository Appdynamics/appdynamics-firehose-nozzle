package sinks

import (
	"fmt"

	"github.com/cloudfoundry/sonde-go/events"
)

type DataPoint struct {
	Metric, Source string
	Value int64
}

type ControllerEventSerializer struct{}

func NewControllerEventSerializer() *ControllerEventSerializer {
	return &ControllerEventSerializer{}
}

func (w *ControllerEventSerializer) BuildHttpStartStopEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func (w *ControllerEventSerializer) BuildLogMessageEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func (w *ControllerEventSerializer) BuildValueMetricEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func (w *ControllerEventSerializer) BuildCounterEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func (w *ControllerEventSerializer) BuildErrorEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func (w *ControllerEventSerializer) BuildContainerEvent(event *events.Envelope) interface{} {
	return genericSerializer(event)
}

func genericSerializer(event *events.Envelope) *DataPoint {
		return &DataPoint{
			Metric: fmt.Sprintf("Server|Component:nozzle-tier|Custom Metrics|%v|%v|%v", event.GetOrigin(), event.GetDeployment(), event.GetIndex()),
			Value: int64(event.GetValueMetric().GetValue()),
			Source: event.GetOrigin()}
}
