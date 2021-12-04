package main

import (
	"fmt"
	"math"
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

// O(n) time | O(1) space
func (m *Metric) Sum() float64 {
	var sum float64
	for _, p := range m.Points {
		sum += p.Value
	}
	return sum
}

// O(n) time | O(1) space
func (m *Metric) Avg() float64 {
	return m.Sum() / float64(len(m.Points))
}

// O(n) time | space O(1)
func (m *Metric) Max() float64 {
	var max float64
	for _, p := range m.Points {
		if max < p.Value {
			max = math.Max(max, p.Value)
		}
	}
	return max
}

// O(n) time | space O(1)
func (m *Metric) Min() float64 {
	min := math.Inf(1)
	for _, p := range m.Points {
		if min > p.Value {
			min = math.Min(min, p.Value)
		}
	}
	return min
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
	fmt.Printf("sum: %f\n", gprcLatencyMetric.Sum())
	fmt.Printf("avg: %f\n", gprcLatencyMetric.Avg())
	fmt.Printf("max: %f\n", gprcLatencyMetric.Max())
	fmt.Printf("min: %f\n", gprcLatencyMetric.Min())
}

/*
sum: 2.200000
avg: 0.733333
max: 1.200000
min: 0.200000
*/
