package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkideal/cli"
	"github.com/ofonimefrancis/spaceship/common"
	"github.com/ofonimefrancis/spaceship/common/config"
	"github.com/ofonimefrancis/spaceship/common/log"
	"github.com/ofonimefrancis/spaceship/common/mgo"
	"github.com/ofonimefrancis/spaceship/features/account"
)

func main() {
	cli.Run(new(config.PackageFlags), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*config.PackageFlags)
		initContext, finishInit := context.WithCancel(context.Background())

		r := gin.Default()

		r.Use(common.EnsureHTTPVersion())
		r.Use(common.SecureHeaders())
		r.Use(common.SilenceSomePanics())

		database := mgo.New(argv.DBHost, argv.DBName)
		r.Use(mgo.DBConnectionMiddleware(database))

		log.Info("Registering account feature...")
		accountHandler := account.NewHandler(initContext, database)
		accountFacade := account.NewHTMLFacade(accountHandler)
		accountFacade.RegisterRoutes(r.Group(account.BasePath))

		finishInit()
		return http.ListenAndServe(fmt.Sprintf(":%d", argv.Port), r)
	})
}
