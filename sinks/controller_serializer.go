package sinks

import (
	"fmt"

	"github.com/cloudfoundry/sonde-go/events"
)

type DataPoint struct {
	Metric, Source string
	Value          int64
}

type ControllerEventSerializer struct {
	tier string
}

func NewControllerEventSerializer(tier_name string) *ControllerEventSerializer {
	return &ControllerEventSerializer{tier: tier_name}
}

func (w *ControllerEventSerializer) BuildHttpStartStopEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func (w *ControllerEventSerializer) BuildLogMessageEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func (w *ControllerEventSerializer) BuildValueMetricEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func (w *ControllerEventSerializer) BuildCounterEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func (w *ControllerEventSerializer) BuildErrorEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func (w *ControllerEventSerializer) BuildContainerEvent(event *events.Envelope) interface{} {
	return genericSerializer(event, w.tier)
}

func genericSerializer(event *events.Envelope, tier string) *DataPoint {
	return &DataPoint{
		Metric: fmt.Sprintf("Server|Component:%v|Custom Metrics|%v|%v|%v", tier, event.GetOrigin(), event.GetDeployment(), event.GetIndex()),
		Value:  int64(event.GetValueMetric().GetValue()),
		Source: event.GetOrigin()}
}
