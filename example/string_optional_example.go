package main

import (
	"github.com/gorpher/optional/v2"
	"log"
)

func main() {
	opt := optional.StringOptional("string")
	if opt.HashError() {
		println(opt.GetError())
	}
	b := opt.String()
	log.Print(b)
	s := opt.ToUpper().IsUpper().HashError()
	log.Print(s)
	slice := optional.SliceStringOptional("good", "nice", "pretty").ToUpper().Slice()
	log.Println(slice)
}
