package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/mcandre/accio"
)

var flagDebug = flag.Bool("debug", false, "Toggle debug mode")
var flagInstall = flag.Bool("install", false, "Install the configured packages")
var flagDestructo = flag.Bool("destructo", false, "Remove all configured executables")
var flagHelp = flag.Bool("help", false, "Show usage information")
var flagVersion = flag.Bool("version", false, "Show version information")

func main() {
	flag.Parse()

	if *flagHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *flagVersion {
		fmt.Println(accio.Version)
		os.Exit(0)
	}

	config, err := accio.Load()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if flagDebug != nil {
		config.Debug = *flagDebug
	}

	if config.Debug {
		log.Printf("Configuration: %v\n", spew.Sdump(config))
	}

	if *flagInstall {
		if err2 := config.Install(); err2 != nil {
			panic(err2)
		}

		os.Exit(0)
	}

	if *flagDestructo {
		if err2 := config.Destructo(); err2 != nil {
			panic(err)
		}

		os.Exit(0)
	}

	flag.PrintDefaults()
	os.Exit(1)
}
