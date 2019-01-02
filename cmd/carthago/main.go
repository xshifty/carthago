package main

import (
	"os"
	"log"

    uuid "github.com/satori/go.uuid"
    shop "github.com/xshifty/carthago"
)

var (
	info = log.New(os.Stdout, "[ INFO ] ", log.LstdFlags)
    warn = log.New(os.Stderr, "[ WARN ] ", log.LstdFlags)
)

func main() {
    info.Println("Welcome to Carthago")

    product, err := shop.NewProduct(uuid.NewV4().String(), "Golang Book", shop.UNIT)
    if err != nil {
        warn.Panicln(err)
    }

    info.Printf("ID: %s, Description: %s, Measure: %v", product.ID, product.Description, product.Unit.String())
}
