package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	model "github.com/xshifty/carthago/model"
	provider "github.com/xshifty/carthago/provider"
)

var (
	info = log.New(os.Stdout, "[ INFO ] ", log.LstdFlags)
	warn = log.New(os.Stderr, "[ WARN ] ", log.LstdFlags)
)

var mux = &sync.Mutex{}

type productRepositoryMock struct {
	data map[model.ID]map[string]interface{}
}

func NewProductRepositoryMock() *productRepositoryMock {
	productRepo := productRepositoryMock{make(map[model.ID]map[string]interface{})}
	for i := 1; i <= 30; i = i + 1 {
		mux.Lock()
		productRepo.data[model.ID(fmt.Sprintf("%04d", i))] = map[string]interface{}{
			"description": fmt.Sprintf("Mock Product #%d", i),
		}
		mux.Unlock()
	}

	return &productRepo
}

func (repo *productRepositoryMock) FindAll() model.ProductList {
	productList := model.ProductList{}

	for id, detail := range repo.data {
		product := model.NewProduct(id, detail["description"].(string))
		productList = append(productList, product)
	}

	return productList
}

func (repo *productRepositoryMock) FindByID(id model.ID) (*model.Product, error) {
	detail, err := repo.data[id]
	if !err {
		return nil, model.ErrProductNotFound
	}

	return model.NewProduct(id, detail["description"].(string)), nil
}

var repositoryFactory = provider.NewRepositoryFactory(provider.RepositoryContainer{
	"productRepository": NewProductRepositoryMock(),
})

func main() {
	info.Println("Welcome to Carthago")

	repo, err := repositoryFactory.Get("productRepository")
	if err != nil {
		warn.Panicln(err)
	}
	productRepository := repo.(*productRepositoryMock)

	for _, product := range productRepository.FindAll() {
		info.Printf("ID: %v (%s)", product.ID(), product.Description())
	}

	prod, err := productRepository.FindByID(model.ID("0003"))
	if err != nil {
		warn.Panicln(err)
	}

	info.Printf("%s (%s)\n", prod.ID(), prod.Description())

	prod, err = productRepository.FindByID(model.ID("1000"))
	if err != nil {
		warn.Fatalln(err)
	}
}
