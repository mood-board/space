package config

import "github.com/mkideal/cli"

type PackageFlags struct {
	cli.Helper
	Test   bool   `cli:"test" usage:"Run in test mode. Affects API sandboxes etc.." dft:"false"`
	Port   int    `cli:"p, port" usage:"Port the app is currently running on" dft:"4000"`
	DBHost string `cli:"db-host" usage:"mongoDB host" dft:"mongo"`
	DBName string `cli:"db-name" usage:"mongoDB name" dft:"opay"`
}
