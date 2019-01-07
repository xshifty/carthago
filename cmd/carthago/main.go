package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
	dic "github.com/xshifty/carthago-dic"
	mock "github.com/xshifty/carthago/mock"
	model "github.com/xshifty/carthago/model"
	provider "github.com/xshifty/carthago/provider"
	types "github.com/xshifty/carthago/types"
)

var (
	info = log.New(os.Stdout, "[ INFO ] ", log.LstdFlags)
	warn = log.New(os.Stderr, "[ WARN ] ", log.LstdFlags)

	container         = dic.New()
	repositoryFactory = provider.NewRepositoryFactory(container)
)

type (
	ProductRepository interface {
		FindAll() model.ProductList
		FindByID(id types.ID) (*model.Product, error)
	}
)

func main() {
	container.Set("db:conn", func(container *dic.Container) interface{} {
		conn, err := sql.Open("postgres", "user=postgres host=localhost sslmode=disable")
		if err != nil {
			panic(err)
		}
		return conn
	})

	container.Set("db:listener", func(container *dic.Container) interface{} {
		return pq.NewListener("", 10*time.Second, time.Minute, func(ev pq.ListenerEventType, err error) {
			if err != nil {
				fmt.Println(err.Error())
			}
		})
	})

	container.Set("repository:Product", func(container *dic.Container) interface{} {
		conn, err := container.Get("db:conn")
		if err != nil {
			panic(err)
		}

		listener, err := container.Get("db:listener")
		if err != nil {
			panic(err)
		}

		return mock.NewProductRepository(
			conn.(*sql.DB),
			listener.(*pq.Listener))
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
}
