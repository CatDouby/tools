package main

import (
	"fmt"

	"github.com/CatDouby/tools/def"
)

type MaxInt int64
type NerPerson struct {
	def.Person
	Phone string
}

func hello() {
	fmt.Println("hello func")
}
