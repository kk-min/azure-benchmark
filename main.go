package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var outputDir = flag.String("o", "./latency_samples/", "Output directory for latency samples (default: ./latency_samples/)")
var burstCount = flag.Int("b", 1, "Number of bursts to send (default: 1)")
var burstSize = flag.Int("s", 1, "Number of requests to send in a burst (default: 1)")
var burstIAT = flag.Int("i", 10000, "Inter-arrival time between bursts in milliseconds (default: 10000)")
var parallel = flag.Bool("p", false, "Send requests in parallel (default: false)")

func main() {
	log.Info("Starting application...")
	flag.Parse()

	outputDirPath := filepath.Join(*outputDir, time.Now().Format(time.RFC850))
	log.Infof("Creating dir for this benchmark at `%s`", outputDirPath)
	if err := os.MkdirAll(outputDirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
