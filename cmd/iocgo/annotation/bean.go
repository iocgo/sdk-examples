package annotation

import (
	"github.com/iocgo/sdk/gen/annotation"
)

type Bean struct {
	*annotation.Anon
	N          string `annotation:"name=name,default="`
	Alias      string `annotation:"name=alias,default="`
	Initialize string `annotation:"name=init,default="`
	Qualifier  string `annotation:"name=qualifier,default="`
}

var _ annotation.M = (*Bean)(nil)

func (g Bean) As() annotation.M {
	return annotation.Inject{
		IsLazy:     false,
		N:          g.N,
		Alias:      g.Alias,
		Initialize: g.Initialize,
		Qualifier:  g.Qualifier,
		Singleton:  false,
		Config:     "yes, i'm golang ~",
	}
}
