package main

import (
	"log"

	"github.com/guni1192/couml/libcouml"
)

func runContainer() {
	c := libcouml.NewContainer()
	workDir := "./rootfs"
	args := []string{"/bin/sh"}

	if err := c.Run(libcouml.NewProcess(args, workDir)); err != nil {
		log.Fatal("Could not container run: ", err)
	}
}
