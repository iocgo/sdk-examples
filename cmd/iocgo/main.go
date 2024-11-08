package main

import (
	"bincooo/sdk-examples/cmd/iocgo/annotation"
	"github.com/iocgo/sdk/gen"
	"github.com/iocgo/sdk/gen/tool"
)

func init() {
	// gin
	gen.Alias[annotation.GetMapping]()
	gen.Alias[annotation.PutMapping]()
	gen.Alias[annotation.DelMapping]()
	gen.Alias[annotation.PostMapping]()

	// ioc
	gen.Alias[annotation.Bean]()

	// cobra
	gen.Alias[annotation.Cobra]()
}

func main() {
	tool.Process()
}
