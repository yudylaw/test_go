package man

import (
	"fmt"
)

type Address struct {
	City   string
	Street string
}

//结构体的内存布局是连续的
//go 中没有类, 结构体就是她的类
type StrongMan struct { // tags
	Weight  int "weight of man"
	Height  int "height of man"
	int         //匿名字段
	Address     //匿名字段, 实现继承
}

//工厂构造方法
//可以私有化 StrongMan ( strongMan ), 强制使用工厂方法
func NewStrongMan(w, h int) *StrongMan {
	//	sm := make(StrongMan) //结构体不支持 make
	//	sm := new(StrongMan)
	//	sm.Weight = w
	//	sm.Height = h
	//	return sm
	return &StrongMan{Weight: w, Height: h} //简洁的初始化方式
}

//初始化函数, 先于 main 方法执行
func init() {
	fmt.Println("package man init.")
}

//有两个原因需要使用指针接收者。
//首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
//其次，方法可以修改接收者指向的值。
//结构体方法
//func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
func (sm *StrongMan) SayHello(height int) {
	sm.Height = height
	fmt.Printf("SayHello w:%d, h:%d \n", sm.Weight, sm.Height)
}

//指针方法和值方法都可以在指针或非指针上被调用
func (sm StrongMan) SayBye(height int) {
	sm.Height = height
	fmt.Printf("SayBye w:%d, h:%d \n", sm.Weight, sm.Height)
}

func (sm StrongMan) Search() {
	fmt.Println("StrongMan search")
}

//实现 Stringers String() 接口
func (sm StrongMan) String() string {
	return fmt.Sprintf("String() weight=%d, height=%d", sm.Weight, sm.Height)
}
