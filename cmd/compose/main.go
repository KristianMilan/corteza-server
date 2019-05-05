package main

import (
	"log"
	"os"

	context "github.com/SentimensRG/ctx"
	"github.com/SentimensRG/ctx/sigctx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/namsral/flag"

	"github.com/crusttech/crust/internal/subscription"
	"github.com/crusttech/crust/internal/version"

	compose "github.com/crusttech/crust/compose"
	system "github.com/crusttech/crust/system"
)

func main() {
	// log to stdout not stderr
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("Starting "+os.Args[0]+", version: %v, built on: %v", version.Version, version.BuildTime)

	ctx := context.AsContext(sigctx.New())

	compose.Flags("compose")
	system.Flags("system")

	subscription.Flags()

	flag.Parse()

	if err := system.Init(ctx); err != nil {
		log.Fatalf("Error initializing: %+v", err)
	}
	if err := compose.Init(ctx); err != nil {
		log.Fatalf("Error initializing: %+v", err)
	}

	var command string
	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	switch command {
	case "help":
		flag.PrintDefaults()
	default:
		// Checks subscription, will os.Exit(1) if there is an error
		// Disabled for now, system service is the only one that validates subscription
		// ctx = subscription.Monitor(ctx)

		if err := compose.StartRestAPI(ctx); err != nil {
			log.Fatalf("Error starting/running: %+v", err)
		}
	}
}