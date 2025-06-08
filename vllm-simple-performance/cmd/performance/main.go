package main

import (
	"log"
	"sync"
	"time"


	"github.com/willschipp/vllm-simple-performance/internal/core" 
	"github.com/willschipp/vllm-simple-performance/internal/util" 
)

func getMetrics(fn func(), interval time.Duration, stop <-chan struct{}) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		case <-stop:
			return
		}
	}
}

func main() {
	log.Println("starting")

	// load the configuration
	config, err := util.LoadConfig()
	if err != nil {
		log.Fatalf("can't load config: %v",err)
	}

	var wg sync.WaitGroup
	url := config.Endpoint.Url
	model := config.Endpoint.Model
	prompt := config.Endpoint.Prompt
	metricUrl := config.Metrics.Url
	metricInterval := config.Metrics.Interval
	metricOutput := config.Metrics.Output

	// start the timer
	stop := make(chan struct{})
	go getMetrics(func() {
		core.GetMetrics(metricUrl,metricOutput)
	},time.Duration(metricInterval), stop)

	for range 10 {
		wg.Add(1)
		go core.SendPrompt(url,model,prompt,&wg)
	}

	wg.Wait()	
	log.Println("All requests completed.")
	//wait 10 seconds more for the metric pull
	time.Sleep(10 * time.Second)
	close(stop) //stop the metric poller
	log.Println("Metric polling complete")
}