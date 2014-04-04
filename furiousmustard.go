package main

import (
	"flag"
	"os/user"
	"path/filepath"
	"time"
	"code.google.com/p/gcfg"
	"github.com/garyburd/redigo/redis"
	"github.com/op/go-logging"
)

// Command ling flags
var verboseFlag = flag.Bool("v", false, "Show verbose debug information")
var configFlag = flag.String("c", "", "Use alternative config file")

// Config
var Config struct {
	Redis struct {
		ConnectionString	string
	}
}

// Logger
var log = logging.MustGetLogger("furiousmustard")

// Redis connection pool
var redisPool = &redis.Pool{
	MaxIdle: 2,
	IdleTimeout: 60 * time.Second,
	Dial: func () (redis.Conn, error) {
		c, err := redis.Dial("tcp", Config.Redis.ConnectionString)
		if err != nil {
			return nil, err
		}
		return c, err
	},
	TestOnBorrow: func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		return err
	},
}

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
