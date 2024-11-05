package annotation

import (
	"github.com/iocgo/sdk/gen/annotation"
)

type PostMapping struct {
	*annotation.Anon
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*PostMapping)(nil)

func (g PostMapping) As() annotation.M {
	return annotation.Router{
		Method: "POST",
		Path:   g.Path,
	}
}
