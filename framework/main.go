package main

import (
	"fmt"
	"service"
)

func main() {
	fmt.Println("Initialize Web Server")
	service := new(service)
	service.start()
}