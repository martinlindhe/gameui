package ui

import (
	"log"
	"path/filepath"
)

// returns full path for assets on disk
func assetPath(s string) string {
	if s != "" && s[0] != '/' && s[0] != '.' {
		return assetPath("../" + s)
	}
	res, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
