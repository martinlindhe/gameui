package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetPath(t *testing.T) {
	tests := map[string]string{
		// input, expect
		".":          `.*[/\\]+martinlindhe[/\\]+gameui$`,
		"./..":       `.*[/\\]+martinlindhe$`,
		"./examples": `.*[/\\]+gameui[/\\]+examples$`,
		"/tmp":       `[/\\]+tmp$`,
	}
	for in, ex := range tests {
		assert.Regexp(t, ex, assetPath(in))
	}
}
