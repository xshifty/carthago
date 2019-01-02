package main

import (
	"os"
	"log"
)

var (
	info = log.New(os.Stdout, "[ INFO ] ", log.LstdFlags)
)

func main() {
    info.Println("Welcome to Carthago")
}
