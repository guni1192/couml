package main

import (
	"log"

	"github.com/guni1192/couml/libcouml"
)

func runContainer() {
	c := libcouml.NewContainer()
	cmd := "/bin/sh"
	workDir := "/home/vagrant/.cromwell/containers/cc3s3izSmKBWwAdj"

	if err := c.Run(libcouml.NewProcess(cmd, workDir)); err != nil {
		log.Fatal("Could not container run: ", err)
	}
}
