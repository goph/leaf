package leaf_test

import (
	"testing"

	"github.com/goph/leaf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindataLoader_Load(t *testing.T) {
	assetFn := func(string) ([]byte, error) {
		return []byte("Hello World!"), nil
	}

	loader := leaf.NewBindataLoader(assetFn)

	template, err := loader.Load("path/to/template.html")
	require.NoError(t, err)

	assert.Equal(t, []byte("Hello World!"), template)
}
