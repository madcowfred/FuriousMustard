package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"github.com/op/go-logging"
	"os/user"
	"path/filepath"
)

// Command ling flags
var configFlag = flag.String("c", "", "Use alternative config file")
var verboseFlag = flag.Bool("v", false, "Show verbose debug information")

// Config
var Config struct {
	Database struct {
		ConnectionString string
	}
}

// Logger
var log = logging.MustGetLogger("furiousmustard")

func main() {
	// Parse command line flags
	flag.Parse()

	// Set up logging
	var format = logging.MustStringFormatter(" %{level: -8s}  %{message}")
	logging.SetFormatter(format)
	if *verboseFlag {
		logging.SetLevel(logging.DEBUG, "furiousmustard")
	} else {
		logging.SetLevel(logging.INFO, "furiousmustard")
	}

	log.Info("FuriousMustard starting")

	// Find config file
	var cfgFile string
	if len(*configFlag) > 0 {
		cfgFile = *configFlag
	} else {
		// Default to user homedir for config file
		u, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		cfgFile = filepath.Join(u.HomeDir, ".furiousmustard.conf")
	}

	// Read config file
	log.Debug("Reading config from %s", cfgFile)

	err := gcfg.ReadFileInto(&Config, cfgFile)
	if err != nil {
		log.Fatal(err)
	}

}
