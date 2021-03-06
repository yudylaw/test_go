package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"person"
	"person/man"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

const (
	MAX int = 100 //常量
	MIN int = 200
)

//自定义类型
type MyUint32 uint32
type String string

//接口
type IAction interface {
	Search()
}

//Go 语言规范定义了接口方法集的调用规则：
//
//类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
//类型 T 的可调用方法集包含接受者为 T 的所有方法
//类型 T 的可调用方法集不包含接受者为 *T 的方法

type IMath interface {
	ADD()
}

//引用指针均可以调用
func (ui MyUint32) Search() {
	fmt.Println("MyUint32 search")
}

//引用指针均可以调用（如果通过对象实例来访问）
//只能指针来调用 ( 如果通过接口来访问，见testAdd）
func (ui *MyUint32) ADD() {
	fmt.Println("MyUint32 add")
}

func makeMyUint32() MyUint32 {
	var my MyUint32 = 10
	fmt.Println("makeMyUint32")
	return my
}

//math *IMath 错误定义, 原因是接口变量中存储的具体值是不可寻址的
func testAdd(math IMath) {
	math.ADD()
}

func testContext() {
	// Pass a context with a timeout to tell a blocking function that it
	// should abandon its work after the timeout elapses.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		//default:
		//	fmt.Println("default...")
	}
}

func main() {
	//	args := os.Args
	//	for index, arg := range args {
	//		fmt.Printf("args, index:%d, arg:%s\n", index, arg)
	//	}
	fmt.Printf("hello goos:%s\n", runtime.GOOS)
	//name := "北京市"
	//has := strings.HasSuffix(name, "市")
	//if has {
	//	index := strings.LastIndex(name, "市")
	//	city := string(name[0:index])
	//	fmt.Println(city)
	//} else {
	//	fmt.Println("has is false")
	//}
	//
	//	fmt.Println(time.Now().Unix())

	//	var str1, str2 string
	//	str1 = getCurrentDir()
	//	str2 = getParentDir(str1)
	//	fmt.Println(str1)
	//	fmt.Println(str2)

	//testChannel()
	//	testPerson();
	//	man.Say();
	//	fmt.Println(MAX);
	//	testFor();
	//	testSwitch();
	//	testDefer()
	//testPointer();
	//testArray()
	//	testMap();
	//	testClosure()
	//	testStruct()
	//	testInterface()
	//	testIO();
	//testHttp()
	//	testRange()
	//	testScan()
	//	testWrite()
	//	testJson()
	//	testErrs()
	//	testPanic()
	//testSelect()
	//testTimer()
	//	testFilePath()
	//testDate()
	//testReg()
	//	testTimer2()
	//	testSplit()
	//	flag := CompareVersion("CA1.4.0_Iphone", "CA1.4.0_Iphone", "Iphone")
	//	fmt.Println("flag=", flag)
	//testContext()
	//	testReflect()
	//testNil()
	//	testSort()
	//testPtr()
	//testTime()
	nums := []int{1, 2, 3}
	nums = nums[:2]
	fmt.Printf("nums=%v", nums)
	fmt.Println("end of main")
}

func testHttp() {
	//	var ui MyUint32 = 10;
	http.Handle("/test", String("hello test"))
	http.Handle("/bug", String("hello bug"))
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":4040", nil) // nil
	if err != nil {
		log.Fatal(err)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	value, err := ioutil.ReadAll(req.Body)
	fmt.Printf("value=%v, err=%v", value, err)
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
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
	ui2 := new(MyUint32)
	//	testAdd(ui) //接受者必须是指针类型
	testAdd(ui2)
	my := makeMyUint32()
	my.ADD()
	if t, ok := action.(MyUint32); ok { //接口类型断言
		fmt.Printf("%T is type of MyUint32\n", t)
	}
	//type-switch 类型判断
	switch t := action.(type) {
	case IAction:
		fmt.Printf("%T is IAction, value:%v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Println("type not found")
	}
	action.Search()
	//只要实现了接口的方法，就满足多态, 没有包依赖
	//	var sm = man.StrongMan{Weight:10, Height:11}
	//	action = sm
	//	action.Search()
	//	fmt.Println(sm)
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

	value, ok := kv["name"] //判断是否存在, 和取值

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
	//常量不可以取地址
	//	fmt.Println(&MAX)
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

	//	runtime.GOMAXPROCS(7) // n - 1

	//通道的方向，通过使用方向注解来限制协程对通道的操作
	//	var send_only chan<- int        // channel can only receive data
	//	var recv_only <-chan int        // channel can only send data

	var myChan chan int = make(chan int)        // 无缓冲
	var signalChan chan int = make(chan int, 2) // cap > 0, 有缓存信道

	//		myChan <- 10;// deadlock，main 线程阻塞在这里, 等待消费者消费

	go write(myChan)
	go read(myChan, signalChan)

	fmt.Println("signal %d", <-signalChan)
}

//通道迭代模式
//for x := range container.Iter() { ... }
func (c *MyUint32) Iter() <-chan int {
	ch := make(chan int)
	return ch
}

func read(ch chan int, sch chan int) {
	for {
		fmt.Println("read %d", <-ch)
		sch <- 0
	}
}

func write(ch chan int) {
	for {
		ch <- 10
		fmt.Println("write")
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

func testScan() {
	//scan
	//	fmt.Println("input your name:")
	//	var name, addr string
	//	fmt.Scanln(&name, &addr) //空格分割
	//	fmt.Printf("name is :%s, addr:%s \n", name, addr)
	//bufio
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("bufio, input something:")

	//	n, err := inputReader.Read(buf) //n 表示读到的字节数
	//	if (n == 0) { break}

	input, err := reader.ReadString('\n')
	if err == nil {
		fmt.Printf("input was:%s\n", input)
	}
}

func testWrite() {
	file, err := os.OpenFile("/home/yudylaw/yudy/test.log", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("failed to OpenFile")
		return
	}
	defer file.Close()
	output := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		output.WriteString("i am yudy\n")
	}
	output.Flush() //bufio 必须有
}

func testJson() {
	add := &man.Address{"合肥", "繁华大道"}
	js, _ := json.Marshal(add)
	fmt.Printf("json is:%s\n", js)
}

func testErrs() {
	err := errors.New("error happened")
	fmt.Println(err)
}

func badCall() {
	panic("bad end")
}

func testPanic() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\n") //执行不到
}

func testSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9) //10^9 ns = 1s

	//	time.Tick(1e9) //定时器通道
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		//功能类型 switch + 事件驱动
		//特点：一次性、阻塞
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}

func testTimer() {
	ticker := time.NewTicker(1 * time.Second)

	go func() {

		for t := range ticker.C {
			fmt.Printf("tick %v\n", t)
		}

	}()

	ch := make(chan int) //无缓冲阻塞队列
	ch <- 0
	fmt.Println("end..")
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDir(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func testFilePath() {
	files, _ := ioutil.ReadDir("E:\\go_workspace\\test_go\\src\\person")
	for _, fi := range files {
		if !fi.IsDir() {
			fmt.Println(fi.Name())
		}
	}
}

func testDate() {
	preDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")

	preTime, _ := time.Parse("2006-01-02", preDate)

	fmt.Println(preTime)

	date, _ := time.Parse("2006-01-02", "2016-05-27")
	fmt.Println(date)

	isBefore := date.Before(preTime)
	fmt.Println(isBefore)
}

func testReg() {
	reg, err := regexp.Compile("advert_([0-9]{4}-[0-9]{2}-[0-9]{2})\\.utf8")
	fmt.Println(err)

	//	match := reg.MatchString("advert_2016-02-12.utf8")
	//	fmt.Println(match)

	fmt.Println(reg.FindStringSubmatch("advert_2016-02-12.utf8"))
}

func testTimer2() {
	startTimer(func() {
		fmt.Println("hello...")
	})
	time.Sleep(100 * time.Second)
}

//每天固定时间点执行
func startTimer(f func()) {
	go func() {
		for {
			now := time.Now()

			next := time.Date(now.Year(), now.Month(), now.Day(), 17, 41, 0, 0, now.Location())

			if next.Before(now) {
				next = next.Add(24 * time.Hour)
			}

			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}

func CompareVersion(sourceCV string, destCV string, OS string) bool {
	//cv例子: CA1.4.0_Iphone, CA1.4.0_Android
	if len(sourceCV) < 14 || len(destCV) < 14 {
		return false
	}

	sourceCV = sourceCV[2:]
	destCV = destCV[2:]

	sourceList := strings.Split(sourceCV, "_")
	destList := strings.Split(destCV, "_")

	if len(sourceList) != 2 || len(destList) != 2 {
		return false
	}

	if sourceList[1] != OS || destList[1] != OS {
		return false
	}

	source := sourceList[0]
	dest := destList[0]

	sourceVec := strings.Split(source, ".")
	destVec := strings.Split(dest, ".")

	if len(sourceVec) != 3 || len(destVec) != 3 {
		return false
	}

	//atoi (表示ascii to integer)
	sourceMajor, _ := strconv.Atoi(sourceVec[0])
	sourceMinor, _ := strconv.Atoi(sourceVec[1])
	sourceBuild, _ := strconv.Atoi(sourceVec[2])

	destMajor, _ := strconv.Atoi(destVec[0])
	destMinor, _ := strconv.Atoi(destVec[1])
	destBuild, _ := strconv.Atoi(destVec[2])

	// 主版本更高
	if sourceMajor > destMajor {
		return true
	}

	// 子版本更高
	if sourceMajor == destMajor && sourceMinor > destMinor {
		return true
	}

	// 特性版本更高
	if sourceMajor == destMajor && sourceMinor == destMinor && sourceBuild >= destBuild {
		return true
	}

	return false
}

func testSplit() {

	var name string = "12.hanzhi.nv"

	index := strings.LastIndex(name, ".")

	fmt.Println(name[index:len(name)])
	fmt.Println(name[0:index])

	var atomic = "name=yudy&age=20&cv=IP_2_2_1"
	values := strings.Split(atomic, "&")
	for _, v := range values {
		pair := strings.Split(v, "=")
		if len(pair) > 1 && pair[0] == "cv" {
			fmt.Println(pair[0], pair[1])
		}
	}
}

func testReflect() {
	var a int = 199
	t := reflect.TypeOf(a)
	fmt.Printf("a.type=%v\n", t)

	m := man.StrongMan{Weight: 100, Height: 200}
	m2 := &man.StrongMan{Weight: 120, Height: 230}
	t1 := reflect.TypeOf(m)
	v1 := reflect.ValueOf(m)
	v2 := reflect.ValueOf(m2)
	fmt.Printf("m.type=%v, kind=%v\n", t1, t1.Kind())
	fmt.Printf("m.value=%v, type=%v\n", v1, v1.Type())
	//引用接收者
	search := v1.MethodByName("Search")
	//没有参数
	search.Call(nil)

	//指针接收者
	sayHello := v2.MethodByName("SayHello")
	height := 10
	v := reflect.ValueOf(height)
	parmas := []reflect.Value{v}
	sayHello.Call(parmas)
}

type data struct{}

func (this *data) Error() string { return "" }

func testNilError() error {
	var p *data = nil
	if p == nil {
		fmt.Printf("testNilError.p is nil\n")
	}
	return p
}

func testNil() {
	//interface{} 包含：类型指针，值指针
	//go语言坑之一
	var val interface{} = nil //nil
	if val == nil {
		fmt.Println("val is nil")
	} else {
		fmt.Println("val is not nil")
	}

	var val2 interface{} = (*interface{})(nil) //not nil
	if val2 == nil {
		fmt.Println("val2 is nil")
	} else {
		fmt.Println("val2 is not nil")
	}

	var p *man.StrongMan = nil //nil
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}

	//类型转换后变成 not nil
	//type到interface和interface之间可能是隐式转换
	//interface{type,data}
	var p2 interface{} = interface{}(p)

	if p2 != nil {
		fmt.Println("p2 is not nil")
	}

	var err error = testNilError() //not nil
	if err == nil {
		fmt.Println("err is nil")
	} else {
		fmt.Println("err is not nil")
	}
	t := reflect.TypeOf(err)
	v := reflect.ValueOf(err)
	fmt.Printf("err.Type=%v, err.Value=%v\n", t, v)
}

//It returns 2*ceil(lg(n+1))
func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func testSort() {
	nums := []int{10, 7, 3, 11, 22, 90}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for _, v := range nums {
		fmt.Println(v)
	}
}

func testPtr() {
	//&符号的意思是对变量取地址
	//*符号的意思是对指针取值
	var a int = 1
	var b *int = &a  //取地址
	var c **int = &b //取地址
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a) //*&可以抵消
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}

func testTime() {
	now := time.Now()
	fmt.Printf("%v", now)
}
