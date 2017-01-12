package main

import (
	"flag"
	"log"
	"os"
	"tes"
	worker "tes/worker"
	"tes/worker/slot"
)

func main() {
	config := worker.Config{}

	var configArg string
	flag.StringVar(&configArg, "config", "", "Config File")
	flag.StringVar(&config.ID, "id", config.ID, "Worker ID")
	flag.StringVar(&config.ServerAddress, "server-address", config.ServerAddress, "Server address")
	flag.StringVar(&config.WorkDir, "workdir", config.WorkDir, "Working Directory")
	flag.IntVar(&config.Timeout, "timeout", config.Timeout, "Timeout in seconds")
	flag.IntVar(&config.NumWorkers, "numworkers", config.NumWorkers, "Worker Count")
	flag.StringVar(&config.LogPath, "logpath", config.LogPath, "File path to write logs to")

	flag.Parse()
	tes.LoadConfigOrExit(configArg, &config)
	start(config)
}

func start(config worker.Config) {

	// TODO Good defaults, configuration, and reusable way to configure logging.
	//      Also, how do we get this to default to /var/log/tes/worker.log
	//      without having file permission problems?
	// Configure logging
	if config.LogPath != "" {
		logFile, err := os.OpenFile(config.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println("Can't open log file")
		} else {
			log.SetOutput(logFile)
		}
	}

	// Create the job engine
	eng, err := worker.NewEngine(config)
	if err != nil {
		log.Printf("Error creating worker engine: %s", err)
		return
	}

	// Configure the slot timeout
	idleTimeout := slot.NoIdleTimeout()
	if config.Timeout != -1 {
		idleTimeout = slot.IdleTimeoutAfterSeconds(config.Timeout)
	}

	// Create the slot pool
	slots := make([]*slot.Slot, config.NumWorkers)
	p := slot.NewPool(slots, idleTimeout)

	// Create the slots
	for i := 0; i < config.NumWorkers; i++ {
		// TODO handle error
		slots[i], _ = slot.NewSlot(config.ID, config.ServerAddress, eng)
	}

	// Start the pool
	p.Start()
}
