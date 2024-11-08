package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/iocgo/sdk"
	"github.com/iocgo/sdk/env"
	"github.com/iocgo/sdk/router"
)

// @Inject(lazy="false", name="ginInitializer")
func Initialized(env *env.Environment) sdk.Initializer {
	return sdk.InitializedWrapper(0, func(container *sdk.Container) (err error) {
		// if !env.Config.GetBool("server.debug") {
		// 	gin.SetMode(gin.ReleaseMode)
		// }
		//
		// engine := gin.Default()
		// beans := sdk.ListInvokeAs[router.Router](container)
		//
		// for _, route := range beans {
		// 	route.Routers(engine)
		// }
		//
		// 初始化的接口不应该阻塞，不然无法执行后续动作。需要阻塞主线程可以使用 container.Run(signals ...os.Signal) 方法
		// fmt.Printf("Listening and serving HTTP on 0.0.0.0:%d\n", port)
		// go func() {
		// 	if err = engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		// 		panic(err)
		// 	}
		// }()
		sdk.ProvideTransient(container, sdk.NameOf[*gin.Engine](), func() (engine *gin.Engine, err error) {
			if !env.Config.GetBool("server.debug") {
				gin.SetMode(gin.ReleaseMode)
			}

			engine = gin.Default()
			beans := sdk.ListInvokeAs[router.Router](container)
			for _, route := range beans {
				route.Routers(engine)
			}

			return
		})
		return
	})
}
