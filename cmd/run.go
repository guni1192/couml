package main

import (
	"log"

	"github.com/guni1192/couml/libcouml"
)

func runContainer(process *libcouml.Process) {
	c := libcouml.NewContainer()

	if err := c.Run(process); err != nil {
		log.Fatal("Could not container run: ", err)
	}
}
