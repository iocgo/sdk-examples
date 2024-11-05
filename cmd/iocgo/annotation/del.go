package annotation

import (
	"github.com/iocgo/sdk/gen/annotation"
)

type DelMapping struct {
	*annotation.Anon
	Path string `annotation:"name=path,default=/"`
}

var _ annotation.M = (*DelMapping)(nil)

func (g DelMapping) As() annotation.M {
	return annotation.Router{
		Method: "DELETE",
		Path:   g.Path,
	}
}
