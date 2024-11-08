package cobra

import (
	"bincooo/sdk-examples/px"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iocgo/sdk"
	"github.com/iocgo/sdk/cobra"
	"github.com/iocgo/sdk/env"
	"github.com/iocgo/sdk/errors"

	_ "github.com/iocgo/sdk"
)

type RootCommand struct {
	container *sdk.Container
	engine    *gin.Engine
	env       *env.Environment

	Name string `cobra:"name" short:"n" usage:"定义名称参数"`
	Age  int    `cobra:"age, per" usage:"定义年龄参数"`
}

// @Cobra(name="rootCobra"
//
//	version = "v1.0.0"
//	use	    = "sdk-examples"
//	short   = "iocgo 工具使用示例"
//	long    = "项目地址: https://www.github.com/iocgo/sdk-examples"
//	run     = "Run"
//
// )
func New(container *sdk.Container, engine *gin.Engine, env *env.Environment, config string) (ic cobra.ICobra, err error) {
	ic = cobra.ICobraWrapper(&RootCommand{
		container: container,
		engine:    engine,
		env:       env,

		Name: "123",
	}, config)
	return
}

func (rc *RootCommand) Run(cmd *cobra.Command, args []string) {
	println("root command running ...  " + rc.Name)
	port := rc.env.Config.GetInt("server.port")
	if port == 0 {
		port = 8080
	}

	if err := Run(rc.container); err != nil {
		panic(err)
	}

	fmt.Printf("Listening and serving HTTP on 0.0.0.0:%d\n", port)
	if err := rc.engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

func Run(container *sdk.Container) (err error) {
	println(container.HealthLogger())
	ctx := errors.New(func(e error) (ok bool) { err = e; return })
	defer ctx.Throw()
	{
		// 自动代理
		bean := errors.Try1(ctx, func() (px.Echo, error) {
			return sdk.InvokeBean[px.Echo](container, "model.A")
		})
		err = bean.Echo("白居易")
		if err != nil {
			println("error: ", err.Error())
			err = nil
		}
	}

	return
}

func warp2[C, T, E any](t T, e E, exec func(T, E) (C, error)) func() (C, error) {
	return func() (C, error) {
		return exec(t, e)
	}
}
