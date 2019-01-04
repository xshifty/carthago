package main

import (
	"log"
	"os"

	mock "github.com/xshifty/carthago/mock"
	model "github.com/xshifty/carthago/model"
	provider "github.com/xshifty/carthago/provider"
	types "github.com/xshifty/carthago/types"
)

var (
	info = log.New(os.Stdout, "[ INFO ] ", log.LstdFlags)
	warn = log.New(os.Stderr, "[ WARN ] ", log.LstdFlags)

	container         = provider.NewContainer()
	repositoryFactory = provider.NewRepositoryFactory(container)
)

type (
	ProductRepository interface {
		FindAll() model.ProductList
		FindByID(id types.ID) (*model.Product, error)
	}
)

func main() {
	container.Add("repository:Product", func(container *provider.Container) interface{} {
		return mock.NewProductRepository()
	})

	info.Println("| Welcome to Carthago")

	repo, err := repositoryFactory.Get("Product")
	if err != nil {
		warn.Panicln(err)
	}
	productRepository := repo.(ProductRepository)

	info.Println("| Printing all products")
	for _, product := range productRepository.FindAll() {
		info.Printf("| ID: %v (%s)", product.ID(), product.Description())
	}

	info.Printf("| Fetching data for product ID: %s", mock.IDTable[3])
	prod, err := productRepository.FindByID(mock.IDTable[3])
	if err != nil {
		warn.Panicln(err)
	}
	info.Printf("| ID: %s (%s)\n", prod.ID(), prod.Description())

	invalidID := types.ID("b169e9b0-e6c8-52c2-9985-8aeb8eb18f03")
	info.Printf("| Fetching data for product ID: %s", invalidID)
	prod, err = productRepository.FindByID(invalidID)
	if err != nil {
		warn.Fatalf("| %s", err)
	}
	info.Printf("| ID: %s (%s)\n", prod.ID(), prod.Description())
}
