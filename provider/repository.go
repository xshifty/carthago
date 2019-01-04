package provider

import "fmt"

type repositoryFactoryHandler struct {
	container *Container
}

func NewRepositoryFactory(container *Container) *repositoryFactoryHandler {
	return &repositoryFactoryHandler{
		container: container,
	}
}

func (factory *repositoryFactoryHandler) Get(repositoryName string) (interface{}, error) {
	repository, err := factory.container.Get(fmt.Sprintf("repository:%s", repositoryName))
	if err != nil {
		return nil, err
	}

	return repository, nil
}
