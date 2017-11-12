package leaf

// Loader is responsible for loading and parsing template content from an arbitrary source.
type Loader interface {
	// Load looks a template up in the underlying source by it's name.
	Load(name string) ([]byte, error)
}

// CachedLoader makes sure that templates only have to be loaded once from the source (eg. file) which makes
// template lookups faster.
type CachedLoader struct {
	loader Loader
	cache  map[string][]byte
}

// NewCachedLoader accepts a loader and wraps it in a CachedLoader.
func NewCachedLoader(loader Loader) *CachedLoader {
	return &CachedLoader{
		loader: loader,
		cache:  make(map[string][]byte),
	}
}

func (l *CachedLoader) Load(name string) ([]byte, error) {
	if template, ok := l.cache[name]; ok {
		return template, nil
	}

	template, err := l.loader.Load(name)
	if err != nil {
		return nil, err
	}

	l.cache[name] = template

	return template, nil
}
