package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"person"
	"person/man"
	"runtime"
	"time"
	"unsafe"
	"reflect"
)

const MAX int = 100 //常量

//自定义类型
type MyUint32 uint32
type String string

//接口
type IAction interface {
	Search()
}

func (ui MyUint32) Search() {
	fmt.Println("MyUint32 search")
}

func main() {
	fmt.Printf("hello goos:%s\n", runtime.GOOS)
	//testChannel();
	//	testPerson();
	//	man.Say();
	//	fmt.Println(MAX);
	//	testFor();
	//	testSwitch();
	//	testDefer()
	//testPointer();
	//	testArray();
	//	testMap();
//	testClosure()
		testStruct()
	//	testInterface();
	//	testIO();
	// testHttp();
	//	testRange()
}

func testHttp() {
	//	var ui MyUint32 = 10;
	http.Handle("/test", String("hello test"))
	http.Handle("/bug", String("hello bug"))
	err := http.ListenAndServe(":4040", nil) // nil
	if err != nil {
		log.Fatal(err)
	}
}

//http Handler interface
func (str String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str)
}

//http Handler interface
func (ui MyUint32) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello http")
}

func testIO() {
	file, err := os.Open("/home/yudylaw/yudy/phpEclipse/eclipse.ini")
	if err != nil {
		log.Fatal("File Not Found.")
	}
	defer file.Close() // defer 延迟关闭

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func testInterface() {
	var ui MyUint32 = 1230
	var action IAction = ui
	action.Search()
	//只要实现了接口的方法，就满足多态, 没有包依赖
	var sm = man.StrongMan{Weight:10, Height:11}
	action = sm
	action.Search()
	fmt.Println(sm)
}

//闭包, 函数也是值
func testClosure() {
	get := sum()
	result := get(10)
	fmt.Println(result)
	result2 := get(7)
	fmt.Println(result2)
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置日志格式
	log.Println("testClosure")
	//	where := func() {
	//	    _, file, line, _ := runtime.Caller(1)
	//	    log.Printf("%s:%d", file, line)
	//	}
	//	where()
}

func sum() func(a int) int {
	tmp := 10 //返回值是函数的闭包函数, 其局部变量的值会被保留
	return func(x int) int {
		tmp *= x
		return tmp
	}
}

func testMap() {
	kvMap := make(map[string]string) //不要用 new 来声明 map
	kvMap["address"] = "BEIJING"
	fmt.Println(kvMap)

	var kv = map[string]string{"age": "xxx"}
	fmt.Println(kv)

	delete(kv, "age") //删除 key
	fmt.Println(kv)

	value, ok := kv["name"] //判断是否存在, 活取值

	fmt.Println("value:%s, ok:%b", value, ok)
}

func testArray() {
	var arr [2]string
	arr[0] = "hello"
	arr[1] = "world"
	fmt.Println(arr)

	//Go 中切片用得更广泛，效率高
	//slice 切片, 切片（slice）是对数组一个连续片段的引用 (长度可变的数组)
	p := []int{133, 2323, 200, 10, 11}
	//	fmt.Println(p, len(p)); //长度
	//	fmt.Println(p[2:2]);//empty
	//	fmt.Println(p[1:2]);//(1,2 - 1]
	fmt.Println(p[:2])
	fmt.Println(p[2:])
	fmt.Println(cap(p))    //容量
	p = append(p, 9, 8, 7) // append
	fmt.Println(cap(p))    //slice 容量自动增长

	for _, v := range p { // v 只是拷贝, 无法修改切片 p
		fmt.Printf("v:%d \n", v)
	}
}

func testDefer() {
	//defer 语句会延迟函数的执行直到上层函数返回。
	defer fmt.Println("defer") //延迟执行, defer 栈，LIFO 先进后出
	fmt.Println("end of the func")
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //保存了当前调用的上下文
	}
	foo()
}

func foo() (age int, name string) {
	//defer 打印函数返回值
	defer func() {
		fmt.Printf("age:%d, name:%s\n", age, name)
	}()
	defer exit(enter("foo")) //trace简单写法, enter 函数没有延迟执行, exit 函数延迟执行
	age = 10
	name = "yudylau"
	fmt.Println("foo end")
	return
}

func enter(str string) string {
	fmt.Println("enter:", str)
	return str
}

func exit(str string) {
	fmt.Println("exit:", str)
}

func testSwitch() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday { // switch { 同一行
	case today + 0:
		fmt.Println("Today.")
		//			break;// break 可以省略
	case today + 1, today + 3: //多值, 表达式
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	i := 1

	switch i {
	case 1:
		fmt.Println("i == 1")
		fallthrough //同时满足下一个 case 条件
	case 2:
		fmt.Println("i == 1 or 2")
	case 3:
		fmt.Println("i == 3")
	}
}

func testFor() {
	//Go 只有 for 一种循环结构，没有 while
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	// while ( for )
	num := 2
	for num < 3 {
		num++
	}
	//	for {
	//		//死循环
	//	}
	fmt.Println(num)
}

func testPointer() {
	var a *int
	if a == nil {
		fmt.Println("a is nil")
	}
	//	var num int = 1010;
	//	var p *int = &num;//指针
	//	fmt.Printf("p:%p \n", p);
	//	fmt.Printf("*p:%d", *p);

	s := "good bye"
	var p *string = &s
	s2 := "yudy"
	fmt.Printf("Here is the pointer p: %p\n", p)
	p = &s2 //改变指针的值, 内存地址
	//    *p = "ciao";//改变指针指向的内容
	fmt.Printf("Here is the pointer p: %p\n", p)
	fmt.Printf("Here is the string *p: %s\n", *p)
	fmt.Printf("Here is the string s: %s\n", s)
}

func testStruct() {
	//struct
	sp := man.StrongMan{Weight: 99, Height: 99}
	sp.City = "HF"
	sp.Street = "TEH"
	fmt.Println(sp)
	sp.SayHello(1001)
	fmt.Println(sp.Address)
	sp.SayBye(1002) // 方法接受者是引用, 无法改变引用指向的值
	fmt.Println(sp)
	sp2 := man.NewStrongMan(110, 111)
	fmt.Printf("memory size:%d byte\n", unsafe.Sizeof(sp2)) //内存大小
	//反射
	type1 := reflect.TypeOf(sp)
	fmt.Println(type1) //结构体类型 man.StrongMan
	fmt.Println(type1.Field(0).Tag)
	type2 := reflect.TypeOf(sp2)
	fmt.Println(type2) // 指针 *man.StrongMan
//	fmt.Println(type2.Field(0).Tag) //不支持 panic: reflect: Field of non-struct type
}

func testType() {
	//类型转换
	var num1 uint32 = 101010
	var num2 uint64 = uint64(num1) //必须显式说明类型
	fmt.Println(num2)
}

func testPerson() {
	var a = 10
	var b = 102
	a, b = person.Swap(a, b)
	fmt.Println("a:%d, b:%d", a, b)
	var max = person.Max(a, b)
	fmt.Println("max:%d", max)
}

func testChannel() {
	var myChan chan int = make(chan int)        // 无缓冲
	var signalChan chan int = make(chan int, 2) // 有缓存信道

	//go channel 生产与消费者必须成对出现
	//	myChan <- 10;// deadlock

	go write(myChan)
	go read(myChan, signalChan)

	fmt.Println("signal %d", <-signalChan)
}

func read(ch chan int, sch chan int) {
	for {
		fmt.Println("read")
		fmt.Println("read %d", <-ch)
		sch <- 0
	}
}

func write(ch chan int) {
	for {
		fmt.Println("write")
		ch <- 10
	}
}

func testRange() {
	str := "你好-abc123"
	for pos, ch := range str {
		fmt.Printf("pos:%d, char:%c, unicode:%U, %X \n", pos, ch, ch, []byte(string(ch)))
		//%U, unicode
		//%X, utf8的16进制
	}
}
