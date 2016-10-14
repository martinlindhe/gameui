package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetPath(t *testing.T) {
	tests := map[string]string{
		// input, expect
		".":     `.*[/\\]+farm[/\\]+ui$`,
		"./..":  `.*[/\\]+farm$`,
		"./dir": `.*[/\\]+farm[/\\]+ui[/\\]+dir$`,
		"/dir":  `[/\\]+dir$`,
	}
	for in, ex := range tests {
		assert.Regexp(t, ex, assetPath(in))
	}
}
