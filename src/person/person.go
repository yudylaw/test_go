package person

import (
	"fmt"
)

//内部方法
func test() {
	fmt.Println("person inner method: test");
}

//方法名大写会被导出
func Swap(a, b int) (int, int){
	return b, a;
}

func Max(a, b int) (max int) {
	//go 不支持三目运算符, 例如: max = a > b ? a : b;
	if (a > b) {
		max = a;
	} else { // } else 必须同一行
		max = b;
	}
	return;
}

