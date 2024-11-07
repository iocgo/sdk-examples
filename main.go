package main

import (
	"bincooo/sdk-examples/model"
	"bincooo/sdk-examples/px"
	"bincooo/sdk-examples/wire"
	"github.com/iocgo/sdk"
	"github.com/iocgo/sdk/errors"
	"syscall"
)

func warp[T any](t T, exec func(T) error) func() error {
	return func() error {
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
	defer ctx.Throw()
	{
		container := sdk.NewContainer()
		errors.Try(ctx, warp(container, wire.Injects))

		sdk.ProvideBean[sdk.Initializer](container, "mainInitializer", Run)
		if err := container.Run(syscall.SIGINT, syscall.SIGTERM); err != nil {
			panic(err)
		}
	}
}

func Run() (sdk.Initializer, error) {
	return sdk.InitializedWrapper(1, func(container *sdk.Container) (err error) {
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

		ctx = errors.New(nil)
		defer ctx.Throw()
		{
			bean := errors.Try1(ctx, warp2[*model.B](container, "model.B", sdk.InvokeBean))
			bean.Echo()
		}

		return
	}), nil
}
