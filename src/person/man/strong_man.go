package man

import (
	"fmt"
)

type StrongMan struct {
	Weight int
	Height int
}

//初始化函数, 先于 main 方法执行
func init() {
	fmt.Println("package man init.")
}

//有两个原因需要使用指针接收者。
//首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
//其次，方法可以修改接收者指向的值。
//结构体方法
func (sm *StrongMan) SayHello(height int) {
	sm.Height = height
	fmt.Printf("SayHello w:%d, h:%d \n", sm.Weight, sm.Height)
}

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
