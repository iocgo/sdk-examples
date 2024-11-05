package main

import (
	"bincooo/sdk-examples/model"
	"bincooo/sdk-examples/px"
	"bincooo/sdk-examples/wire"
	"github.com/iocgo/sdk"
	"github.com/iocgo/sdk/errors"
	"github.com/iocgo/sdk/proxy"
	"syscall"
)

func warp[T any](t T, exec func(T) error) func() error {
	return func() error {
		return exec(t)
	}
}

func warp1[T any](t T, exec func(T) (T, error)) func() (T, error) {
	return func() (T, error) {
		return exec(t)
	}
}

func warp2[C, T, E any](t T, e E, exec func(T, E) (C, error)) func() (C, error) {
	return func() (C, error) {
		return exec(t, e)
	}
}

func main() {
	ctx := errors.New(nil)
	defer ctx.Do()

	container := sdk.NewContainer()
	errors.Try(ctx, warp(container, wire.Injects))

	sdk.ProvideBean[sdk.Initializer](container, "mainInitializer", Run)
	if err := container.Run(syscall.SIGINT, syscall.SIGTERM); err != nil {
		panic(err)
	}
}

func Run() (sdk.Initializer, error) {
	return sdk.InitializedWrapper(1, func(container *sdk.Container) (err error) {
		println(container.HealthLogger())
		// 直接用代理接口获取实例
		// bean, err := sdk.InvokeAs[px.Echo](container, "model.A")
		// if err != nil {
		// 	panic(err)
		// }
		//
		// err = bean.Echo("白居易")
		// if err != nil {
		// 	println("error: ", err.Error())
		// }

		ctx := errors.New(func(e error) (ok bool) { err = e; return })
		{
			defer ctx.Do()
			// 通过实例构建代理
			bean := errors.Try1(ctx, warp2[*model.A](container, "model.A", sdk.InvokeBean))
			echo := errors.Try1(ctx, warp1[px.Echo](bean, proxy.New))
			err = echo.Echo("白居易")
			if err != nil {
				println("error: ", err.Error())
				err = nil
			}
		}

		ctx = errors.New(nil)
		{
			defer ctx.Do()
			bean := errors.Try1(ctx, warp2[*model.B](container, "model.B", sdk.InvokeBean))
			bean.Echo()
		}

		return
	}), nil
}
