package main

import (
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var (
	reMovie  = regexp.MustCompile(`^(.+) \[.*?\][\[\(](\d+)[\)\]]$`)
	scanning = false
)

func FileScanner() {
	// Call scanFiles at the start
	scanFiles()

	// Call scanFiles once every 60 seconds thereafter
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			scanFiles()
		}
	}
}

func scanFiles() {
	if scanning == true {
		return
	}

	scanning = true
	defer func() {
		scanning = false
	}()

	// Scan movie directories
	for _, moviePath := range Config.Paths.Movies {
		moviePath, err := filepath.EvalSymlinks(moviePath)
		if err != nil {
			log.Error(err.Error())
			return
		}

		err = filepath.Walk(moviePath, visitMovies)
		if err != nil {
			log.Error(err.Error())
			return
		}
	}

	scanning = false
}

func visitMovies(path string, f os.FileInfo, err error) error {
	log.Debug("Visited: %s", path)
	return nil
}
