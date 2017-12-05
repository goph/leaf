# Leaf

[![Build Status](https://img.shields.io/travis/goph/leaf.svg?style=flat-square)](https://travis-ci.org/goph/leaf)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/leaf?style=flat-square)](https://goreportcard.com/report/github.com/goph/leaf)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/leaf)

**The missing template container for Go.**

Go comes with a great template engine, proven to work well in many applications
(Kubernetes Helm, Consul templates, etc). Although in simple cases it's rather enjoyable to work with it,
"traditional" web development can quickly become cumbersome when using a lot of templates.

Features like the idea of template inheritance, a central template container, sophisticated template loading
mechanisms and many other are all found in well-known template engines commonly used for web development.
Unfortunately these are all missing from Go, which is understandable since [simplicity](https://talks.golang.org/2015/simplicity-is-complicated.slide#4)
is one of the key values of the Go language. But we still need a solution for these problems.

Leaf tries to address some of the issues mentioned above.
It provides a loading mechanism for templates as well as the ability to define custom loading mechanisms.
Once the templates are loaded, they are compiled into an immutable container.
Preferably this is done when the application starts, so you can fail early if something is wrong.

It also provides a way to reload templates when they are stored in files, so you can develop your application without
recompiling it on every template change.


## Known limitations

- Defined templates cannot depend on each other for the moment.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
