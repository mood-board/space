package main

import (
	"fmt"
	"time"

	"github.com/mkideal/cli"
	"github.com/ofonimefrancis/spaceship/common/config"
)

func main() {
	cli.Run(new(config.PackageFlags), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*config.PackageFlags)
		fmt.Println(argv)
		fmt.Println(time.Now().Format(time.RFC3339))
		return nil
	})
}
