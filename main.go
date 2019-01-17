package main

import (
	"fmt"

	"github.com/mkideal/cli"
	"github.com/ofonimefrancis/spaceship/common/config"
)

func main() {
	cli.Run(new(config.PackageFlags), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*config.PackageFlags)
		fmt.Println(argv)
		return nil
	})
}
