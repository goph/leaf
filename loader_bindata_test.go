package leaf_test

import (
	"bufio"
	"bytes"
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

	tpl, err := loader.Load("path/to/template.html")
	require.NoError(t, err)

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	err = tpl.Execute(writer, nil)
	require.NoError(t, err)
	writer.Flush()

	assert.Equal(t, "Hello World!", buf.String())
}
