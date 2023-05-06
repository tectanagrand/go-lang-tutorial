package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

type Person struct {
	Name
}