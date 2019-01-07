package provider

import (
	"fmt"

	dic "github.com/xshifty/carthago-dic"
)

type repositoryFactoryHandler struct {
	container *dic.Container
}

func NewRepositoryFactory(container *dic.Container) *repositoryFactoryHandler {
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
