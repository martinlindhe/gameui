package ui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetPath(t *testing.T) {
	assert.Regexp(t, ".*/farm/ui$", assetPath(""))
	assert.Regexp(t, ".*/farm/ui$", assetPath("."))
	assert.Regexp(t, ".*/farm$", assetPath("./.."))
	assert.Regexp(t, ".*/farm/ui/dir", assetPath("./dir"))
	assert.Regexp(t, "^/dir$", assetPath("/dir"))
}
