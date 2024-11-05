package annotation

import (
	"github.com/iocgo/sdk/gen/annotation"
)

type PutMapping struct {
	*annotation.Anon
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*PutMapping)(nil)

func (g PutMapping) As() annotation.M {
	return annotation.Router{
		Method: "PUT",
		Path:   g.Path,
	}
}
