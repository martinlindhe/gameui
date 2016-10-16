package ui

import (
	"log"
	"os"
	"path/filepath"
)

// returns full path for assets on disk
func assetPath(s string) string {
	res, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}

	// file exists?
	if _, err := os.Stat(res); err == nil {
		return res
	}
	log.Fatal("path not found", s)
	return s
}
