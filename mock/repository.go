package mock

import (
	"errors"

	"database/sql"

	"github.com/lib/pq"

	model "github.com/xshifty/carthago/model"
	types "github.com/xshifty/carthago/types"
)

var ErrProductNotFound = errors.New("Product not found.")

type productRepositoryMock struct {
	db       *sql.DB
	listener *pq.Listener
}

func NewProductRepository(db *sql.DB, listener *pq.Listener) *productRepositoryMock {
	productRepo := productRepositoryMock{
		db:       db,
		listener: listener,
	}

	return &productRepo
}

func (repo *productRepositoryMock) FindAll() model.ProductList {
	rows, err := repo.db.Query("SELECT id, description FROM carthago.product")
	if err != nil {
		panic(err)
	}

	var (
		productList model.ProductList
		id          string
		description string
	)

	for {
		ok := rows.Next()
		if !ok {
			break
		}

		rows.Scan(&id, &description)
		productList = model.ProductList(append(productList, model.NewProduct(types.ID(id), description)))
	}

	return productList
}

func (repo *productRepositoryMock) FindByID(id types.ID) (*model.Product, error) {
	rows, err := repo.db.Query("SELECT id, description FROM carthago.product WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	var description string

	if rows.Next() {
		rows.Scan(&id, &description)
		return model.NewProduct(id, description), nil
	}

	return nil, errors.New("No data found.")
}
