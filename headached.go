package main

import (
	"encoding/json"
	"github.com/oniichaNj/headached/lib/corrupt"
	"github.com/oniichaNj/headached/lib/entropyexhaustion"
	"github.com/oniichaNj/headached/lib/load"
	"log"
	"os"
	"sync"
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
	/* Should we increase load average? */
	EnableCPUSpike bool
	/* The interval of seconds to wait between CPU usage */
	MinCPUSpikeSeconds int
	MaxCPUSpikeSeconds int
	/* The duration of a CPU spike, in seconds */
	CPUSpikeDuration int
	/* Should we exhaust system entropy? */
	EnableEntropyExhaustion bool
}

func main() {

	/*
	 * Logging to STDERR is neater because we can then let something else deal with rotation
	 * and choose what file to put things in.
	 */

	errLog := log.New(os.Stderr, "", log.Ldate|log.Ltime)

	/* Open and set up the configuration file. */
	cfgf, err := os.Open("/etc/headached.json")
	if err != nil {
		errLog.Fatal(err)
	}
	defer cfgf.Close()

	decoder := json.NewDecoder(cfgf)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		errLog.Fatal(err)
	}

	var wg sync.WaitGroup

	/*
	 * Load the components from lib/ that we want to use.
	 * If you wrote your own components, add them here.
	 */

	if config.EnableCorruption {
		wg.Add(1)
		go corrupt.Init(config.MinCorruptSeconds, config.MaxCorruptSeconds, config.CorruptDirs, config.CorruptNbytes, errLog)
	}

	if config.EnableCPUSpike {
		wg.Add(1)
		go load.Init(config.MinCPUSpikeSeconds, config.MaxCPUSpikeSeconds, config.CPUSpikeDuration, errLog)
	}

	if config.EnableEntropyExhaustion {
		wg.Add(1)
		go entropyexhaustion.Init(errLog)
	}

	wg.Wait()

}
