package config

import "github.com/mkideal/cli"

type PackageFlags struct {
	cli.Helper
	Port int `cli:"p, port" usage:"Port the app is currently running on"`
}
