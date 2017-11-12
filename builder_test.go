package leaf_test

import (
	"bufio"
	"bytes"
	"html/template"
	"testing"

	"github.com/goph/leaf"
	"github.com/goph/leaf/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBuilder_Define(t *testing.T) {
	loader := new(mocks.Loader)

	loader.On("Load", "path/to/base.html").Return([]byte(`{{ block "header" . }}{{ end }}{{ block "main" . }}{{ end }}{{ block "footer" . }}{{ end }}`), nil).Once()
	loader.On("Load", "path/to/header.html").Return([]byte(`{{ define "header" }}Header{{ end }}`), nil).Once()
	loader.On("Load", "path/to/footer.html").Return([]byte(`{{ define "footer" }}Footer{{ end }}`), nil).Once()
	loader.On("Load", "path/to/main.html").Return([]byte(`{{ define "main" }}Main{{ end }}`), nil).Once()

	builder := leaf.NewBuilder(loader)

	builder.Define("my_template", "path/to/base.html", "path/to/header.html", "path/to/footer.html", "path/to/main.html")

	tpl, err := builder.Get("my_template")
	require.NoError(t, err)

	buf := new(bytes.Buffer)
	writer := bufio.NewWriter(buf)
	err = tpl.Execute(writer, nil)
	require.NoError(t, err)
	writer.Flush()

	assert.Equal(t, "HeaderMainFooter", buf.String())
}

func TestBuilder_Set(t *testing.T) {
	expectedTemplate, err := template.New("base").Parse(`Hello World!`)
	require.NoError(t, err)

	builder := leaf.NewBuilder(nil)

	builder.Set("my_template", expectedTemplate)

	actualTemplate, err := builder.Get("my_template")
	require.NoError(t, err)

	assert.Equal(t, expectedTemplate, actualTemplate)
}

func TestBuilder_Compile(t *testing.T) {
	expectedTemplate, err := template.New("base").Parse(`Hello World!`)
	require.NoError(t, err)

	builder := leaf.NewBuilder(nil)

	builder.Set("my_template", expectedTemplate)

	container, err := builder.Compile()
	require.NoError(t, err)

	actualTemplate, err := container.Get("my_template")
	require.NoError(t, err)

	assert.Equal(t, expectedTemplate, actualTemplate)
}
