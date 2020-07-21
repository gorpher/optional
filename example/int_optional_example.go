package main

import (
	"github.com/gorpher/optional/v2"
	"log"
)

func main() {
	b := optional.BoolOptional(true)
	var v bool
	log.Println(b.String())
	log.Println(b.Bool())
	log.Println(b.Int())
	err := b.SetBool(&v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)
}
