package provider

import (
	"errors"
)

type RepositoryContainer map[string]interface{}

var (
	ErrContainerKeyNotFound = errors.New("Container key not found.")
)

type repositoryFactoryHandler struct {
	container RepositoryContainer
}

func NewRepositoryFactory(container RepositoryContainer) *repositoryFactoryHandler {
	return &repositoryFactoryHandler{
		container: container,
	}
}

func (factory *repositoryFactoryHandler) Get(key string) (interface{}, error) {
	repository, err := factory.container[key]
	if !err {
		return nil, ErrContainerKeyNotFound
	}

	return repository, nil
}
