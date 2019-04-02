package main

import (
	"log"

	"github.com/guni1192/couml/libcouml"
)

func runContainer() {
	c := libcouml.NewContainer()
	if err := c.Run(); err != nil {
		log.Fatal("Could not container run: ", err)
	}
}
