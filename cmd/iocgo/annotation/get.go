package annotation

import (
	"github.com/iocgo/sdk/gen/annotation"
)

type GetMapping struct {
	*annotation.Anon
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*GetMapping)(nil)

func (g GetMapping) As() annotation.M {
	return annotation.Router{
		Method: "GET",
		Path:   g.Path,
	}
}
