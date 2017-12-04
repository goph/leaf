package leaf

import (
	"html/template"
	"io"
)

// Builder provides a mutable interface for a template container.
//
// It also implements some additional template loading strategies (like expressing dependencies between templates).
//
// In production, the Builder should be compiled into an immutable container when the program starts, so that errors are
// caught early.
type Builder struct {
	loader      Loader
	definitions map[string]Definition
}

// NewBuilder accepts a loader and returns a new Builder.
func NewBuilder(loader Loader) *Builder {
	return &Builder{
		loader:      loader,
		definitions: make(map[string]Definition),
	}
}

func (b *Builder) Get(name string) (*template.Template, error) {
	definition, ok := b.definitions[name]
	if !ok {
		return nil, ErrTemplateNotFound
	}

	return definition.Resolve()
}

func (b *Builder) Execute(name string, wr io.Writer, data interface{}) error {
	tpl, err := b.Get(name)
	if err != nil {
		return err
	}

	return tpl.Execute(wr, data)
}

// Define defines a new template based on dependencies.
func (b *Builder) Define(name string, templates ...string) {
	b.definitions[name] = &loaderDefinition{
		name:      name,
		templates: templates,

		loader: b.loader,
	}
}

// Set allows setting an already created template in the builder.
func (b *Builder) Set(name string, tpl *template.Template) {
	b.definitions[name] = &templateDefinition{
		tpl: tpl,
	}
}

// Compile parses and loads all templates and creates an immutable container from them.
//
// This method should be called early so that template issues do not occur runtime.
func (b *Builder) Compile() (Container, error) {
	templates := make(map[string]*template.Template, len(b.definitions))

	for name, definition := range b.definitions {
		t, err := definition.Resolve()
		if err != nil {
			return nil, err
		}

		templates[name] = t
	}

	return &container{templates}, nil
}

// Definition is the building block for the container builder.
//
// Definitions are used internally to resolve templates.
// Template resolution occurs either during "compilation" time (see Compile method of Builder)
// or every time a template is fetched from the Builder. This is necessary, because the Builder is mutable,
// so between two subsequent fetcher calls a template definition for the same name might change.
//
// This behaviour is also useful for template development: templates can be reloaded (eg. when they are files)
// without restarting the application. (Note that this is highly unrecommended for production usage)
type Definition interface {
	// Resolve tries to resolve a template from the definition.
	Resolve() (*template.Template, error)
}

// templateDefinition always resolves to the same template reference.
type templateDefinition struct {
	tpl *template.Template
}

func (d *templateDefinition) Resolve() (*template.Template, error) {
	return d.tpl, nil
}

// loaderDefinition loads a template along with it's dependencies using a user configured template loader.
type loaderDefinition struct {
	name      string
	base      string
	templates []string
	loader    Loader
}

func (d *loaderDefinition) Resolve() (*template.Template, error) {
	tpl := template.New(d.name)

	for _, dependency := range d.templates {
		templateContent, err := d.loader.Load(dependency)
		if err != nil {
			return nil, err
		}

		tpl, err = tpl.Parse(string(templateContent))
		if err != nil {
			return nil, err
		}
	}

	return tpl, nil
}
