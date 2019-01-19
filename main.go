package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mkideal/cli"
	"github.com/ofonimefrancis/spaceship/common/config"
	"github.com/ofonimefrancis/spaceship/common/mgo"
)

func main() {
	cli.Run(new(config.PackageFlags), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*config.PackageFlags)
		fmt.Println(argv)
		fmt.Println(time.Now().Format(time.RFC3339))

		initContext, finishInit := context.WithCancel(context.Background())

		r := gin.Default()

		// r.Use(common.EnsureHTTPVersion())
		// r.Use(common.HeaderCheck())
		// r.Use(common.SecureHeaders())
		// r.Use(common.CacheHeaders(static.BasePath))

		// r.Use(common.SilenceSomePanics())

		database := mgo.New(argv.DBHost, argv.DBName)
		r.Use(mgo.DBConnectionMiddleware(database))

		finishInit()
		return http.ListenAndServe(fmt.Sprintf(":%s", argv.Port), r)
	})
}
