package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	/* The directory to remove files from: */
	CorruptDir string
	/* The interval of seconds to wait between file corruption: */
	MinCorruptSleepSeconds int
	MaxCorruptSleepSeconds int
	/* The interval of seconds to wait between CPU usage */
	MinCPUSpikeSeconds int
	MaxCPUSpikeSeconds int
	/* The duration of a CPU spike, in seconds */
	CPUSpikeDuration int
}

func main() {
	/* Set up a neat log. Should change path to something /var/log before release. */
	e, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}
	defer e.Close()
	errLog := log.New(e, ">>>", log.Ldate|log.Ltime)

	/* Open and set up the configuration file. Should change path to /etc/headached.json before release. */
	cfgf, err := os.Open("headached.json")
	if err != nil {
		errLog.Println(err)
	}
	decoder := json.NewDecoder(cfgf)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		errLog.Println(err)
	}

	for {

	}
}
