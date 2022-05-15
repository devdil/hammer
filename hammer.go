package main

import (
	"flag"
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/montanaflynn/stats"
	"log"
	"net/http"
	"sync"
	"time"
)

type result struct {
	statusCode int
	timeTook   int
}

/*



 */

func main() {
	colorGreen := "\033[32m"
	rps := flag.Int("rps", 1, "Requests per second. Usage rps=1 where 1 is a positive integer")
	dur := flag.Int("dur", 1, "Duration in seconds dur=1 where is 1 is a valid duration in seconds")
	flag.Parse()

	requestsPerSeconds := *rps
	duration := *dur

	var wg sync.WaitGroup
	fmt.Println("                          \\`. \n.--------------.___________) \\\n|//////////////|___________[ ]\n`--------------'           ) (\n                           '-'")
	fmt.Println("\n█░█ ▄▀█ █▀▄▀█ █▀▄▀█ █▀▀ █▀█\n█▀█ █▀█ █░▀░█ █░▀░█ ██▄ █▀▄")
	fmt.Println("Hammer Load Test Tool ")
	fmt.Println("---------------------------------")
	fmt.Println("Requests Per Second: ", requestsPerSeconds)
	fmt.Println("Duration: ", duration)
	fmt.Println("---------------------------------")
	fmt.Println("Hammering.....")
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	s.Color("green")
	s.Start()
	startTimeSeconds := time.Now().Second()

	// create a channel of n requests
	responses := make(chan result, requestsPerSeconds*duration)

	// loop until we have made all the requests
	for {
		currentTimeSeconds := time.Now().Second()
		secondsSinceStarted := currentTimeSeconds - startTimeSeconds
		if secondsSinceStarted == duration {
			break
		}

		for i := 1; i <= requestsPerSeconds; i++ {
			wg.Add(1)
			go GetHomePage("http://diljitpr.net", responses, &wg)
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("\n Collecting results...")
	wg.Wait()
	close(responses)
	s.Stop()
	// iterate through responses and print the result
	var responseTimes []float64
	for result := range responses {
		responseTimes = append(responseTimes, float64(result.timeTook))
	}
	fmt.Println(string(colorGreen))
	fmt.Println("---------------------------------")
	fmt.Println("Summary Statistics               ")
	fmt.Println("---------------------------------")
	d := stats.LoadRawData(responseTimes)
	p99, _ := stats.Percentile(responseTimes, 99)
	p95, _ := stats.Percentile(responseTimes, 95)
	p90, _ := stats.Percentile(responseTimes, 90)
	max, _ := stats.Max(d)
	min, _ := stats.Min(d)
	fmt.Println("P99(ms) : ", p99)
	fmt.Println("P95(ms) : ", p95)
	fmt.Println("P90(ms) : ", p90)
	fmt.Println("Max(ms) : ", max)
	fmt.Println("Min(ms) : ", min)
	fmt.Println("---------------------------------")
}

func GetHomePage(url string, responses chan result, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now().UnixMilli()
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Something went wrong while fetching ", err)
	}
	defer resp.Body.Close()
	timeItTook := time.Now().UnixMilli() - startTime
	res := result{statusCode: resp.StatusCode, timeTook: int(timeItTook)}
	responses <- res
}
