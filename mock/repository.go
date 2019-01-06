package mock

import (
	"errors"
	"fmt"
	"time"

	"database/sql"

	"github.com/lib/pq"

	model "github.com/xshifty/carthago/model"
	types "github.com/xshifty/carthago/types"
)

var (
	ErrProductNotFound = errors.New("Product not found.")
	IDTable            = make(map[int]types.ID)
)

type productRepositoryMock struct {
	//data     map[types.ID]map[string]interface{}
	db       *sql.DB
	listener *pq.Listener
}

func NewProductRepository() *productRepositoryMock {
	db, err := sql.Open("postgres", "user=postgres host=localhost sslmode=disable")
	if err != nil {
		panic(err)
	}

	productRepo := productRepositoryMock{
		db: db,
		listener: pq.NewListener("", 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
			if err != nil {
				fmt.Println(err.Error())
			}
		}),
	}

	return &productRepo
}

func (repo *productRepositoryMock) FindAll() model.ProductList {
	rows, err := repo.db.Query("SELECT id, description, metadata FROM carthago.product")
	if err != nil {
		panic(err)
	}

	var (
		productList model.ProductList
		id          sql.NullString
		description sql.NullString
		metadata    sql.NullString
	)

	for {
		ok := rows.Next()

		if !ok {
			break
		}

		rows.Scan(&id, &description, &metadata)
		fmt.Println(id, description, metadata)
	}

	return productList
}

func (repo *productRepositoryMock) FindByID(id types.ID) (*model.Product, error) {
	/*
		detail, err := repo.data[id]
		if !err {
			return nil, ErrProductNotFound
		}
	*/

	//return model.NewProduct(id, detail["description"].(string)), nil
	return model.NewProduct(id, "hello world"), nil
}
