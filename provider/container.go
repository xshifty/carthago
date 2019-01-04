package provider

import (
	"errors"
)

var (
	ErrDependencyAlreadyExists = errors.New("Dependency already exists.")
	ErrContainerKeyNotFound    = errors.New("Container key not found.")
)

type containerEntry struct {
	key    string
	loader func(*Container) interface{}
	cache  interface{}
	loaded bool
}

func newContainerEntry(key string, loader func(*Container) interface{}) *containerEntry {
	return &containerEntry{
		key:    key,
		loader: loader,
		cache:  nil,
		loaded: false,
	}
}

type Container struct {
	entries map[string]*containerEntry
}

func NewContainer() *Container {
	return &Container{
		entries: make(map[string]*containerEntry),
	}
}

func (container *Container) Add(key string, loader func(*Container) interface{}) error {
	_, err := container.entries[key]
	if err {
		return ErrDependencyAlreadyExists
	}

	container.entries[key] = newContainerEntry(key, loader)

	return nil
}

func (container *Container) Get(key string) (interface{}, error) {
	dep, err := container.entries[key]

	if !err {
		return nil, ErrContainerKeyNotFound
	}

	if !dep.loaded {
		dep.cache = dep.loader(container)
		dep.loaded = true
	}

	return dep.cache, nil
}
