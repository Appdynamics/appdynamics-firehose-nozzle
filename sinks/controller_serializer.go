package sinks

import (
	"fmt"

	"github.com/cloudfoundry/sonde-go/events"
)

type DataPoint struct {
	Metric  string
	Value   int64
	Allowed bool
}

type ControllerEventSerializer struct {
	tier string
}

func NewControllerEventSerializer(tier_name string) *ControllerEventSerializer {
	return &ControllerEventSerializer{tier: tier_name}
}

func (w *ControllerEventSerializer) BuildHttpStartStopEvent(event *events.Envelope) interface{} {
	allowed := false
	return &DataPoint{Metric: "", Value: int64(0), Allowed: allowed}
}

func (w *ControllerEventSerializer) BuildLogMessageEvent(event *events.Envelope) interface{} {
	allowed := false
	return &DataPoint{Metric: "", Value: int64(0), Allowed: allowed}
}

func (w *ControllerEventSerializer) BuildValueMetricEvent(event *events.Envelope) interface{} {
	origin, name, deployment, index := event.GetOrigin(), event.GetValueMetric().GetName(), event.GetDeployment(), event.GetIndex()
	alias, present := FilterMetrics(origin, name)
	prefix := fmt.Sprintf("Server|Component:%v|Custom Metrics", w.tier)
	if present {
		return &DataPoint{
			Metric:  fmt.Sprintf("%v|%v|%v|%v|%v", prefix, alias, name, deployment, index),
			Value:   int64(event.GetValueMetric().GetValue()),
			Allowed: present}
	} else {
		return &DataPoint{Metric: "", Value: int64(0), Allowed: present}
	}
}

func (w *ControllerEventSerializer) BuildCounterEvent(event *events.Envelope) interface{} {
	origin, name, deployment, index := event.GetOrigin(), event.GetCounterEvent().GetName(), event.GetDeployment(), event.GetIndex()
	alias, present := FilterMetrics(origin, name)
	prefix := fmt.Sprintf("Server|Component:%v|Custom Metrics", w.tier)
	if present {
		return &DataPoint{
			Metric:  fmt.Sprintf("%v|%v|%v|%v|%v", prefix, alias, name, deployment, index),
			Value:   int64(event.GetCounterEvent().GetDelta()),
			Allowed: present}
	} else {
		return &DataPoint{Metric: "", Value: int64(0), Allowed: present}
	}
}

func (w *ControllerEventSerializer) BuildErrorEvent(event *events.Envelope) interface{} {
	allowed := false
	return &DataPoint{Metric: "", Value: int64(0), Allowed: allowed}
}

func (w *ControllerEventSerializer) BuildContainerEvent(event *events.Envelope) interface{} {
	allowed := false
	return &DataPoint{Metric: "", Value: int64(0), Allowed: allowed}
}

func FilterMetrics(eventOrigin, eventName string) (string, bool) {
	filters, present := MetricFilter[eventOrigin]
	if present {
		alias := MetricAlias[eventOrigin]
		for _, allowedMetric := range filters {
			if allowedMetric == eventName {
				return alias, true
			}
		}
	}
	return "", false
}
