package leaf_test

import (
	"testing"

	"github.com/goph/leaf"
	"github.com/goph/leaf/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCachedLoader_Load(t *testing.T) {
	wrappedLoader := new(mocks.Loader)

	expectedTemplate := []byte("Hello World!")
	wrappedLoader.On("Load", "path/to/template.html").Return(expectedTemplate, nil).Once()

	loader := leaf.NewCachedLoader(wrappedLoader)

	firstTemplate, err := loader.Load("path/to/template.html")
	require.NoError(t, err)
	assert.Equal(t, expectedTemplate, firstTemplate)

	// Test subsequent call does not call the wrapper again
	secondTemplate, err := loader.Load("path/to/template.html")
	require.NoError(t, err)
	assert.Equal(t, firstTemplate, secondTemplate)
}

func TestCachedLoader_Load_Error(t *testing.T) {
	wrappedLoader := new(mocks.Loader)

	wrappedLoader.On("Load", "path/to/template.html").Return(nil, leaf.ErrTemplateNotFound).Once()

	expectedTemplate := []byte("Hello World!")
	wrappedLoader.On("Load", "path/to/template.html").Return(expectedTemplate, nil).Once()

	loader := leaf.NewCachedLoader(wrappedLoader)

	_, err := loader.Load("path/to/template.html")
	require.Error(t, err)
	assert.Equal(t, leaf.ErrTemplateNotFound, err)

	// Errors are not cached
	actualTemplate, err := loader.Load("path/to/template.html")
	require.NoError(t, err)
	assert.Equal(t, expectedTemplate, actualTemplate)
}
