package main

import (
	"github.com/kai1/osx-activated/internal/app"
	"github.com/kai1/osx-activated/pkg/version"
)

func init() {
	version.Version.Print()
}

func main() {
	if err := app.App(); err != nil {
		panic(err)
	}
}
