package leaf

import (
	"errors"
	"html/template"
	"io"
)

// ErrTemplateNotFound is returned when a template is not found.
var ErrTemplateNotFound = errors.New("template not found")

// Container describes an immutable interface for a structure holding a list of templates.
type Container interface {
	// Get tries to fetch a template from the container.
	//
	// Returns an error when the template cannot be found or cannot be loaded in case of a dynamic container.
	Get(name string) (*template.Template, error)

	// Execute fetches and executes the template.
	//
	// Returns an error when the template cannot be found or cannot be loaded in case of a dynamic container.
	// If the execution ends up in an error, it's returned as well.
	Execute(name string, wr io.Writer, data interface{}) error
}

// container is the default, immutable implementation of Container returned by Builder.Compile.
type container struct {
	templates map[string]*template.Template
}

func (c *container) Get(name string) (*template.Template, error) {
	t, ok := c.templates[name]
	if !ok {
		return nil, ErrTemplateNotFound
	}

	return t, nil
}

func (c *container) Execute(name string, wr io.Writer, data interface{}) error {
	tpl, err := c.Get(name)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, data)
}
