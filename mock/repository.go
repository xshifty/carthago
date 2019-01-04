package mock

import (
	"errors"
	"fmt"

	uuid "github.com/satori/go.uuid"
	model "github.com/xshifty/carthago/model"
	types "github.com/xshifty/carthago/types"
)

var (
	ErrProductNotFound = errors.New("Product not found.")

	IDTable = make(map[int]types.ID)
)

type productRepositoryMock struct {
	data map[types.ID]map[string]interface{}
}

func NewProductRepository() *productRepositoryMock {
	productRepo := productRepositoryMock{make(map[types.ID]map[string]interface{})}
	for i := 0; i < 20; i = i + 1 {
		IDTable[i] = types.ID(uuid.NewV4().String())
		productRepo.data[IDTable[i]] = map[string]interface{}{
			"description": fmt.Sprintf("Mock Product #%d", i+1),
		}
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

func (repo *productRepositoryMock) FindByID(id types.ID) (*model.Product, error) {
	detail, err := repo.data[id]
	if !err {
		return nil, ErrProductNotFound
	}

	return model.NewProduct(id, detail["description"].(string)), nil
}
