package px

import (
	"errors"
	"fmt"
	"github.com/iocgo/sdk/proxy"
	"reflect"
	"strings"
	// ------>>> 代理接口中未被使用的包需要导入 <<<-----
)

type Echo interface {
	Echo(name string) error
}

func ValueType(t any) string {
	ox := reflect.ValueOf(t)
	if ox.Kind() == reflect.Ptr {
		ox = ox.Elem()
	}

	value := fmt.Sprintf("%T", t)
	// 通过静态生成的代理
	if value[0] == '*' && strings.HasSuffix(value, "_px__") {
		field := ox.FieldByName("proto")
		value = field.Type().String()
	}

	if !strings.Contains(value, "/") {
		if value[0] == '*' {
			return "*bincooo/sdk-examples/" + value[1:]
		}
		return "bincooo/sdk-examples/" + value
	}

	return value
}

// @Proxy(target="px.Echo")
func EchoInvocationHandler(ctx *proxy.Context[Echo]) {
	fmt.Println("开始代理...")
	fmt.Println("原型: ", ValueType(ctx.Receiver))

	name := ctx.In[0].(string)
	fmt.Println("入参: ", name)
	name = "李白"
	fmt.Println("修改入参: ", name)
	ctx.In[0] = name

	// 反射修改实例
	// obj := ctx.Receiver
	// ox := reflect.ValueOf(obj)
	// if ox.Kind() == reflect.Ptr {
	// 	ox = ox.Elem()
	// }
	//
	// field := ox.FieldByName("Num")
	//
	// i := field.Int()
	// fmt.Println("Num: ", i)
	// // 非指针类型无法设置值
	// if field.CanSet() {
	// 	field.SetInt(2)
	// }

	ctx.Do()
	// debugger 栈追踪
	// if true {
	// 	panic("testing")
	// }

	var err error

	// 指针类型
	if o := ctx.Out[0]; o != nil {
		err = o.(error)
	}

	fmt.Println("出参: ", err)
	err = errors.New("panic")
	ctx.Out[0] = err
	fmt.Println("修改出参: ", err)

	fmt.Println("结束代理...")
}
