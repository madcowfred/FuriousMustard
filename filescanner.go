package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"
)

var (
	reMovie  = regexp.MustCompile(`^(.+) (?:\[.*?\]|)[\[\(](\d+)[\)\]]\.(?:avi|mkv|mp4)$`)
	scanning = false
)

func FileScanner() {
	scanFiles()

	// Call scanFiles once every 60 seconds
	ticker := time.NewTicker(60 * time.Second)
	for {
		select {
		case <-ticker.C:
			scanFiles()
		}
	}
}

func scanFiles() {
	if (scanning == true) {
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
			log.Warn("EvalSymlinks error: %s", err.Error())
			return
		}

		err = filepath.Walk(moviePath, visitMovies)
		if err != nil {
			log.Warn("Walk error: %s", err.Error())
			return
		}
	}
}

func visitMovies(path string, f os.FileInfo, err error) error {
	// Get a Redis connection from our pool
	conn := redisPool.Get()
	defer conn.Close()

	// Get a filename without the path and match against our scary regexp
	base := filepath.Base(path)
	matches := reMovie.FindStringSubmatch(base)

	// If it matches, do some stuff
	if len(matches) > 0 {
		log.Debug("%s -> %q", base, matches)

		// Check mediainfo
		cmd := exec.Command("mediainfo", "--output=XML", path)
		out, err := cmd.Output()
		if err != nil {
			log.Warn("mediainfo error: %s", err.Error())
			return nil
		}

		log.Debug(string(out))
	}

	return nil
}
