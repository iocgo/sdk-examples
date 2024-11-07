package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iocgo/sdk"
	"github.com/iocgo/sdk/env"
	"github.com/iocgo/sdk/router"
)

// @Inject(lazy="false", name="ginInitializer")
func Initialized(env *env.Environment) sdk.Initializer {
	return sdk.InitializedWrapper(0, func(container *sdk.Container) (err error) {
		port := env.Config.GetInt("server.port")
		if port == 0 {
			return
		}

		if !env.Config.GetBool("server.debug") {
			gin.SetMode(gin.ReleaseMode)
		}

		engine := gin.Default()
		beans := sdk.ListInvokeAs[router.Router](container)
		if err != nil {
			return
		}

		for _, route := range beans {
			route.Routers(engine)
		}

		// 初始化的接口不应该阻塞，不然无法执行后续动作。需要阻塞主线程可以使用 container.Run(signals ...os.Signal) 方法
		fmt.Printf("Listening and serving HTTP on 0.0.0.0:%d\n", port)
		go func() {
			if err = engine.Run(fmt.Sprintf(":%d", port)); err != nil {
				panic(err)
			}
		}()
		return
	})
}
