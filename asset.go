package ui

import (
	"log"
	"os"
	"path/filepath"
)

// returns full path for assets on disk
func assetPath(s string) string {
	if exists(s) {
		res, err := filepath.Abs(s)
		if err != nil {
			log.Fatal(err)
		}
		return res
	}
	if s[0] != '/' && s[0] != '.' {
		return assetPath("../" + s)
	}
	return "INVALID-PATH"
}

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
