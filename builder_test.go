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

	base, err := template.New("base").Parse(`{{ block "header" . }}{{ end }}{{ block "main" . }}{{ end }}{{ block "footer" . }}{{ end }}`)
	require.NoError(t, err)
	loader.On("Load", "path/to/base.html").Return(base, nil).Once()

	header, err := template.New("header").Parse(`{{ define "header" }}Header{{ end }}`)
	require.NoError(t, err)
	loader.On("Load", "path/to/header.html").Return(header, nil).Once()

	main, err := template.New("main").Parse(`{{ define "main" }}Main{{ end }}`)
	require.NoError(t, err)
	loader.On("Load", "path/to/main.html").Return(main, nil).Once()

	footer, err := template.New("footer").Parse(`{{ define "footer" }}Footer{{ end }}`)
	require.NoError(t, err)
	loader.On("Load", "path/to/footer.html").Return(footer, nil).Once()

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
