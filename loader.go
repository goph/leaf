package leaf

import "html/template"

// Loader is responsible for loading and parsing template content from an arbitrary source.
type Loader interface {
	// Load looks a template up in the underlying source by it's name.
	Load(name string) (*template.Template, error)
}

// CachedLoader makes sure that templates only have to be loaded once from the source (eg. file) which makes
// template lookups faster.
type CachedLoader struct {
	loader Loader
	cache map[string]*template.Template
}

// NewCachedLoader accepts a loader and wraps it in a CachedLoader.
func NewCachedLoader(loader Loader) *CachedLoader {
	return &CachedLoader{
		loader: loader,
		cache: make(map[string]*template.Template),
	}
}

func (l *CachedLoader) Load(name string) (*template.Template, error) {
	if tpl, ok := l.cache[name]; ok {
		return tpl, nil
	}

	tpl, err := l.loader.Load(name)
	if err != nil {
		return nil, err
	}

	l.cache[name] = tpl

	return tpl, nil
}
