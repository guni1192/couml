package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"

	"github.com/guni1192/couml/libcouml"
	"github.com/urfave/cli"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		runtime.GOMAXPROCS(1)
		runtime.LockOSThread()

		// TODO: loag config from config.json
		// var spec specs.Spec = LoadConfigJson(path)
		containerConfig := libcouml.ContainerConfig{
			Cwd:  "./rootfs",
			Args: []string{"/bin/sh"},
		}

		libcouml.PrepareRootfs(&containerConfig)
		runContainer()
	}
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "run",
			Usage: "run container",
			Action: func(c *cli.Context) error {
				cmd := exec.Command(os.Args[0], "init")

				cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

				cmd.SysProcAttr = &syscall.SysProcAttr{
					Cloneflags: syscall.CLONE_NEWUSER | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS,
					UidMappings: []syscall.SysProcIDMap{
						{
							ContainerID: 0,
							HostID:      syscall.Getuid(),
							Size:        1,
						},
					},
					GidMappings: []syscall.SysProcIDMap{
						{
							ContainerID: 0,
							HostID:      syscall.Getgid(),
							Size:        1,
						},
					},
				}
				return cmd.Run()
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "init",
					Usage: "init process",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
