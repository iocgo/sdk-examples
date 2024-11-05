package wire

import (
	"github.com/iocgo/sdk"

	_ "bincooo/sdk-examples/gin"
	_ "bincooo/sdk-examples/gin/handler"
	_ "bincooo/sdk-examples/model"
)

// @Gen()
func Injects(container *sdk.Container) error {
	panic("auto implements")
}
