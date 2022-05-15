package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

const COLOR_GREEN string = "\033[32m"

type GenericReporter interface {
	GenerateReport()
}

type GenericReporterGenerator struct {
	Values []float64
}

func (g *GenericReporterGenerator) PrintGeneralStatistics(values []float64) {
	fmt.Println("---------------------------------")
	fmt.Println("Summary Statistics               ")
	fmt.Println("---------------------------------")
	d := stats.LoadRawData(values)
	p99, _ := stats.Percentile(values, 99)
	p95, _ := stats.Percentile(values, 95)
	p90, _ := stats.Percentile(values, 90)
	max, _ := stats.Max(d)
	min, _ := stats.Min(d)
	fmt.Println("P99(ms) : ", p99)
	fmt.Println("P95(ms) : ", p95)
	fmt.Println("P90(ms) : ", p90)
	fmt.Println("Max(ms) : ", max)
	fmt.Println("Min(ms) : ", min)
	fmt.Println("---------------------------------")
}

type HttpReportGenerator struct {
	responses       chan httpResponse
	genericReporter GenericReporterGenerator
}

func (httpReport HttpReportGenerator) GenerateReport() {

	// iterate through responses and print the result
	var responseTimes []float64
	for result := range httpReport.responses {
		responseTimes = append(responseTimes, float64(result.timeTook))
	}
	fmt.Println(string(COLOR_GREEN))
	// call generic reporter
	httpReport.genericReporter.Values = responseTimes
	httpReport.genericReporter.PrintGeneralStatistics(responseTimes)

}
