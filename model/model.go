package model

import (
	"fmt"

	// ------>>> 代理接口中未被使用的包需要导入 <<<-----
	_ "bincooo/sdk-examples/px"
	_ "github.com/iocgo/sdk"
	_ "github.com/iocgo/sdk/proxy"
)

type A struct {
	Num int
}

type B struct {
	*A
}

// @Inject(name="model.A", proxy="bincooo/sdk-examples/px.Echo")
func NewA() *A {
	return &A{}
}

// @Inject(alias="model.B", qualifier="[0]:model.A")
func NewB(a *A) *B {
	return &B{a}
}

func (a A) Echo(name string) error {
	fmt.Printf("%d, A.Echo(%s)\n", a.Num, name)
	return nil
}

func (B) Echo() {
	fmt.Println("B.Echo()")
}
