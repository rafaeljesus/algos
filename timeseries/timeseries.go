package main

import (
	"fmt"
	"sort"
	"time"
)

// p1 : 1.0, 08/10/2021
type Point struct {
	Value     float64
	Timestamp time.Time
}

/*
metric name: ”process.cpu”
labels: {host=”foobar”, app_name=”crouton”, kind=”user”, …}
points: (value=1.7, start=T1, limit=T2)
	(value=0.8, start=T3, limit=T4)
	(value=3.1, start=T5, limit=T6)
	(value=1.1, start=T7, limit=T8)
	(value=0.5, start=T9, limit=T10)
metric name: ”grpc.client.latencies”
labels: {host=”foobar”, app_name=”crouton”, operation=”/GetProjectConfig”, …}
points: (value=, start=T1, limit=T2)
*/
type Metric struct {
	Name   string
	Labels map[string]string
	Points []*Point
}

func (m *Metric) Sum() float64 {
	var sum float64
	for _, p := range m.Points {
		sum += p.Value
	}
	return sum
}

func (m *Metric) Avg() float64 {
	return m.Sum() / float64(len(m.Points))
}

func (m *Metric) Max() float64 {
	var values []float64
	for _, p := range m.Points {
		values = append(values, p.Value)
	}
	sort.Float64s(values)
	return values[len(values)-1]
}

func main() {
	now := time.Now()
	gprcLatencyMetric := &Metric{
		Name: "grpc.client.latencies",
		Labels: map[string]string{
			"app":       "crouton",
			"operation": "/GetProjectConfig",
		},
		Points: []*Point{
			&Point{Value: 0.8, Timestamp: now},
			&Point{Value: 1.2, Timestamp: now},
			&Point{Value: 0.2, Timestamp: now},
		},
	}
	fmt.Println(gprcLatencyMetric.Sum())
	fmt.Println(gprcLatencyMetric.Avg())
	fmt.Println(gprcLatencyMetric.Max())
}
