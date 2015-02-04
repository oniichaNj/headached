package main

import (
	"encoding/json"
	"fmt"
	"github.com/oniichaNj/headached/lib/corrupt"
	"github.com/oniichaNj/headached/lib/entropyexhaustion"
	"github.com/oniichaNj/headached/lib/load"
	"log"
	"os"
	"time"
)

type Config struct {
	/* Should we corrupt files? */
	EnableCorruption bool
	/* The directory to remove files from: */
	CorruptDirs []string
	/* The interval of seconds to wait between file corruption: */
	MinCorruptSeconds int
	MaxCorruptSeconds int
	/* The amount of bytes to corrupt. */
	CorruptNbytes int
	/* The interval of seconds to wait between CPU usage */
	MinCPUSpikeSeconds int
	MaxCPUSpikeSeconds int
	/* The duration of a CPU spike, in seconds */
	CPUSpikeDuration int
	/* Should we exhaust system entropy? */
	EnableEntropyExhaustion bool
}

func main() {

	/* Set up a neat log. Should change path to something /var/log before release. */
	e, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer e.Close()
	errLog := log.New(e, "", log.Ldate|log.Ltime)

	/* Open and set up the configuration file. Should change path to /etc/headached.json before release. */
	cfgf, err := os.Open("headached.json")
	if err != nil {
		errLog.Println(err)
	}
	defer cfgf.Close()
	decoder := json.NewDecoder(cfgf)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		errLog.Println(err)
	}

	/*
	 * Load the components from lib/ that we want to use.
	 * If you wrote your own components, add them here.
	 */

	if config.EnableCorruption {
		go corrupt.Init(config.MinCorruptSeconds, config.MaxCorruptSeconds, config.CorruptDirs, config.CorruptNbytes, errLog)
	}
	go load.Init(config.MinCPUSpikeSeconds, config.MaxCPUSpikeSeconds, config.CPUSpikeDuration, errLog)
	if config.EnableEntropyExhaustion {
		go entropyexhaustion.Init(errLog)
	}

	for {
		time.Sleep(5 * time.Second)

	}
}
