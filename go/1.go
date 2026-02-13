// package main

// import (
// 	"fmt"
// )

//	func main() {
//	   var stockcode=123;
//	   var enddate="203";
//	   var url="code=%d&endDate=%s";
//	   fmt.Printf(url,stockcode,enddate)
//	}
// package main

// import "fmt"

// func main(){
//    var a string = "Runoob"
//    fmt.Println(a)

//    var b,c int=1,2
//    fmt.Println(b,c)
// }

// package main
// import "fmt"
// func main(){
//    var a ="runoob"
//    fmt.Println(a)

//    var b int
//    fmt.Println(b)

//	   var c bool
//	   fmt.Print(c)
//	}
// package main

// import "fmt"
// var x,y int
//  var (
//    a int
//    b bool
//  )
//  var c,d int =1,2
//  var e,f =123,"hello"

//  var g,h=123,"hello"

//	func main(){
//	  fmt.Println(x,y,a,b,c,d,e,f,g,h)
//	}
// package main

// import "fmt"
//
//	func main(){
//	   var a string = "abc"
//	   fmt.Println("hello,world",a)
//	}

// package main

// import "fmt"
// func main(){
//    _,munb,strs := numbers()
//    fmt.Println(munb,strs)
// }
// func numbers()(int,int,string){
//    a,b,c := 1,2,"str"
//    return a,b,c
// }

// package main
// import "fmt"
//
//	func main(){
//	   const leght int = 10
//	   const width int =5
//	   var area int
//	   const a,b,c =1,false,"str"
//	   area = leght * width
//	   fmt.Println("面积为:%d",area)
//	   println()
//	   println(a,b,c)
//	}
// package main

// import "unsafe"
// const(
//
//	a = "abc"
//	b = len(a)
//	c = unsafe.Sizeof(a)
//
// )
//
//	func main(){
//	   println(a,b,c)
//	}
// package main

// import "fmt"
//
//	func main(){
//	   const(
//	      a = iota
//	      b
//	      c
//	      d="ha"
//	      e
//	      f=100
//	      g
//	      h=iota
//	      i
//	   )
//	   fmt.Println(a,b,c,d,e,f,g,h,i)
//	}
// package main

// import "fmt"
// const(
//
//	i=1<<iota
//	j=3<<iota
//	k
//	l
//
// )
//
//	func main(){
//	   fmt.Println("i=",i)
//	   fmt.Println("j=",j)
//	   fmt.Println("k=",k)
//	   fmt.Println("l=",l)
//	}
// package main

// import "fmt"
// func main(){
//    var a bool =true
//    var b bool =false
//    if(a&&b){
//       fmt.Println("true\n")
//    }
//    if(a||b){
//       fmt.Println("true\n")
//    }
//    // a=false
//    // b=true
//    if(a&&b){
//       fmt.Println("true\n")
//    }else{
//       fmt.Println("false\n")
//    }
//    if(!(a&&b)){
//       fmt.Println("true\n")
//    }
// }

// package main

// import "fmt"
// func main(){
//    var a uint = 60
//    var b uint = 13
//    var c uint =0

//    c=b>>2
//    fmt.Printf("%d\n",c)

//    c=a & b
//    fmt.Printf("%d\n",c)
// }

// package main

// import "fmt"
// func main(){
//    var a int =4
//    var ptr *int
//    ptr = &a
//    fmt.Printf("%d\n",*ptr)
// }

// package main

// import "fmt"
// func main(){
// 	var a int=100
// 	var b int=200
// 	var ret int

// 	ret = max(a,b)
// 	fmt.Println("\n",ret)
// }
// func max(num1,num2 int)int{
// 	var result int

// 	if(num1>num2){
// 		result = num1
// 	} else{
// 		result = num2
// 	}
// 	return result
// }
// package main

// import "fmt"

// func swap(x, y string) (string, string) {
//    return y, x
// }

//	func main() {
//	   a, b := swap("Google", "Runoob")
//	   fmt.Println(a, b)
//	}
// package main

// import "fmt"
// func main(){
// 	var a,b,c int
// 	a=10
// 	b=20
// 	c=a+b
// 	fmt.Printf("a=%d,b=%d and c=%d\n",a,b,c)
// }

// package main

// import "fmt"
// func main(){
// 	var n[10]int
// 	var i,j int
// 	for i = 0; i < 10; i++ {
// 		n[i] = i +100
// 	}
// 	for j = 0; j < 10; j++ {
// 		fmt.Printf("[%d]=%d\n",j,n[j])
// 	}
// }

// package main

// import "fmt"
// func main(){
// 	var i,j,k int
// 	balance :=[5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
// 	for i =0;i<5;i++{
// 		fmt.Printf("[%d]=%f\n",i,balance[i])
// 	}
// 	balance2 :=[...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
// 	for j =0;j<5;j++{
// 		fmt.Printf("[%d]=%f\n",j,balance2[j])
// 	}
// 	balance3 :=[5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
// 	for k =0;k<5;k++{
// 		fmt.Printf("[%d]=%f\n",k,balance3[k])
// 	}
// }

// package main

// import "fmt"
//
//	func main(){
//		var a int=10
//		fmt.Printf("%x\n",&a)
//	}

// package main

// import "fmt"
// func main(){
// 	var a int=20
// 	var ip *int

// 	ip=&a
// 	fmt.Printf("%x\n",&a)
// 	fmt.Printf("%x\n",ip)
// 	fmt.Printf("%d\n",*ip)
// }

// 切片
// package main

// import "fmt"
// func main(){
// 	var numbers = make([]int,3,5)
// 	printSlice(numbers)
// }
// func printSlice(x []int){
// 	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
// }

// package main

// import "fmt"
// func main(){
// 	numbers :=[]int{0,1,2,3,4,5,6,7,8}
// 	printSlice(numbers)

// 	fmt.Println("numbers ==",numbers)

// 	fmt.Println("numbers[1:4] ==",numbers[1:4])

// 	fmt.Println("numbers[:3] ==",numbers[:3])

// 	fmt.Println("numbers[4:] ==",numbers[4:])

// 	numbers1 := make([]int,0,5)
// 	printSlice(numbers1)

// 	numbers2 := numbers[:2]
// 	printSlice(numbers2)

// 	numbers3 := numbers[2:5]
// 	printSlice(numbers3)
// }
// func printSlice(x []int){
// 	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
// }

// package main

// import "fmt"

// func main() {
//    var numbers []int
//    printSlice(numbers)

//    /* 允许追加空切片 */
//    numbers = append(numbers, 0)
//    printSlice(numbers)

//    /* 向切片添加一个元素 */
//    numbers = append(numbers, 1)
//    printSlice(numbers)

//    /* 同时添加多个元素 */
//    numbers = append(numbers, 2,3,4)
//    printSlice(numbers)

//    /* 创建切片 numbers1 是之前切片的两倍容量*/
//    numbers1 := make([]int, len(numbers), (cap(numbers))*2)

//    /* 拷贝 numbers 的内容到 numbers1 */
//    copy(numbers1,numbers)
//    printSlice(numbers1)
// }

// func printSlice(x []int){
//    fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
// }

// package main

// import "fmt"

// func main(){
// 	for i,c := range "hello"{
// 		fmt.Printf("index: %d, char:%c\n",i,c)
// 	}
// }

// package main

// import "fmt"

// func main(){
// 	a:=[]int{1,2,3,4,5}
// 	b:=[]int{5,4,3,2,1}
// 	for k,v :=range a {
// 		fmt.Println(k,v)
// 	}
// 	for j,n :=range b{
// 		fmt.Println(j,n)
// 	}
// }

// package main

// import "fmt"
// func main(){
// 	map1 :=make(map[int]float32)

// 	map1[1] =1.0

// 	for key,value :=range map1{
// 		fmt.Printf("key is: %d - value is:%f\n",key,value)
// 	}

// 	for key := range map1{
// 		fmt.Printf("key is:%d\n",key)
// 	}
// 	for _,value :=range map1{
// 		fmt.Printf("value is :%f\n",value)
// 	}
// }

// package main

// import "fmt"
// func produced(ch chan int){
// 	for i :=0; i < 5;i++{
// 		ch <-i
// 		fmt.Println("Produced:",i)
// 	}
// 	close(ch)
// }

// func consumer (ch chan int){
// 	for{
// 		val,ok := <-ch
// 		if !ok {
// 			break
// 		}
// 		fmt.Println("Consumer:",val)
// 	}
// }

// func main(){
// 	ch :=make(chan int,4)
// 	go produced(ch)
// 	consumer(ch)
// }

// package main

// import "fmt"
// func main(){
// 	nums :=[]int{2,3,4}

// 	for _,num :=range nums{
// 		fmt.Println("value:",num)
// 	}
// 	for i := range nums{
// 		fmt.Println("index:",i)
// 	}
// }

// package main

// import "fmt"
// func main(){
// 	intSlice :=[]int{1,2,3,4,5,6}
// 	for index,value :=range	intSlice{
// 		fmt.Println("index:",index,"value:",value)
// 	}

// 	testString :="hello"
// 	for index,value :=range testString{
// 		fmt.Println(index,value)
// 	}
// 	kvs :=map[string]string{"a":"apple","b":"banana"}
// 	for k,v :=range	kvs{
// 		fmt.Printf("%s=%s\n",k,v)
// 	}
// 	for i,c :=range "golang"{
// 		fmt.Println(i,c)
// 	}
// }

// package main

// import "fmt"
// func main(){
// 	var siteMap map[string]string
// 	siteMap = make(map[string]string)

// 	siteMap ["google"] = "1"
// 	for site :=range siteMap{
// 		fmt.Println(site,"为",siteMap[site])
// 	}
// 	name,ok :=siteMap["Facebook"]
// 	if(ok){
// 		fmt.Println("Facebook",name)
// 	} else{
// 		fmt.Println("Facebook 无")
// 	}

// }

// package main

// import "fmt"
// func main(){
// 	countryCapitalMap :=map[string]string{"France":"Paris","Italy":"Rome"}
// 	fmt.Println("地图")

// 	for country :=range countryCapitalMap{
// 		fmt.Println(country,"首都",countryCapitalMap[country])
// 	}
// 	// 删除
// 	delete(countryCapitalMap,"France")
// 	for country :=range countryCapitalMap{
// 		fmt.Println(country,"首都",countryCapitalMap[country])
// 	}
// }

// package main

// import "fmt"
// func factorial(n int)int{
// 	if n == 0{
// 		return 1
// 	}
// 	return n * factorial(n-1)
// }
// func main(){
// 	fmt.Println(factorial(10))
// }

// package main

// import "fmt"
// func fibonacci(n int)int{
// 	if n<2{
// 		return n
// 	}
// 	return fibonacci(n-2) + fibonacci(n-1)
// }
// func main(){
// 	var i int
// 	for i = 0;i<11;i++{
// 		fmt.Printf("%d\t",fibonacci(i))
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// func sqrtRecursive(x, guess, prevGuess, epsilon float64) float64 {
//         if diff := guess*guess - x; diff < epsilon && -diff < epsilon {
//                 return guess
//         }

//         newGuess := (guess + x/guess) / 2
//         if newGuess == prevGuess {
//                 return guess
//         }

//         return sqrtRecursive(x, newGuess, guess, epsilon)
// }

// func sqrt(x float64) float64 {
//         return sqrtRecursive(x, 1.0, 0.0, 1e-9)
// }

// func main() {
//         x := 25.0
//         result := sqrt(x)
//         fmt.Printf("%.2f 的平方根为 %.6f\n", x, result)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main(){
// 	str :="123"
// 	num,err :=strconv.Atoi(str)
// 	if err !=nil{
// 		fmt.Printf(":",err)
// 	} else{
// 		fmt.Printf("'%s' :%d\n",str,num)
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main(){
// 	num :=123
// 	str :=strconv.Itoa(num)
// 	fmt.Printf("%d : '%s'\n",num,str)
// }

// package main

// import (
// 	"fmt"
// 	"strconv"
// )
// func main(){
// 	str :="3.14"
// 	num,err :=strconv.ParseFloat(str,64)
// 	if err !=nil{
// 		fmt.Printf(":",err)
// 	} else{
// 		fmt.Printf("%f :'%s'\n",num,str)
// 	}
// }

// package main

// import "fmt"
// func main(){
// 	var i interface{}="hello"
// 	str,ok :=i.(string)
// 	if ok{
// 		fmt.Printf("'%s' is a string\n",str)
// 	} else{
// 		fmt.Println("conversion failed")
// 	}
// }

// package main

// import "fmt"
// type Writer interface{
// 	Writer([]byte)(int,error)
// }
// type StringWriter struct{
// 	str string
// }

// func (sw *StringWriter) Writer(data []byte)(int,error){
// 	sw.str +=string(data)
// 	return len(data),nil
// }
// func main(){
// 	var w Writer = &StringWriter{}
// 	sw :=w.(*StringWriter)
// 	sw.str = "hello"

// 	fmt.Printf(sw.str)
// }

// package main

// import "fmt"
// func printValue(v interface{}){
// 	switch v :=v.(type){
// 	case int:
// 		fmt.Println("Integer:",v)
// 	case string:
// 		fmt.Println("String:",v)
// 	default:
// 		fmt.Printf("Unknown type")
// 	}
// }
// func main(){
// 	printValue(42)
// 	printValue("hello")
// 	printValue(3.14)
// }

// package main

// import (
// 	"fmt"
// )
// func printValue(val interface{}){
// 	fmt.Printf("Value:%v,type:%t\n",val,val)
// }
// func main(){
// 	printValue(42)
// 	printValue("hello")
// 	printValue(3.14)
// 	printValue([]int{1,2})
// }

// package main

// import "fmt"

// func main(){
// 	var i interface{}=42
// 	if str,ok := i.(string);ok{
// 		fmt.Println("String:",str)
// 	} else{
// 		fmt.Println("no")
// 	}
// }

// package main

// import "fmt"
// func printType(val interface{}){
// 	switch v :=val.(type){
// 	case int:
// 		fmt.Println("Integer:",v)
// 	case string:
// 		fmt.Println("String:",v)
// 	case float64:
// 		fmt.Println("Float:",v)
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }
// func main(){
// 	printType(42)
// 	printType("hello")
// 	printType(3.14)
// 	printType([]int{1,2,3})
// }

// package main

// import (
// 	"fmt"
// )

// type Phone interface {
//     call()
// }

// type NokiaPhone struct {
// }

// func (nokiaPhone NokiaPhone) call() {
//     fmt.Println("I am Nokia, I can call you!")
// }

// type IPhone struct {
// }

// func (iPhone IPhone) call() {
//     fmt.Println("I am iPhone, I can call you!")
// }

// func main() {
//     var phone Phone

//     phone = new(NokiaPhone)
//     phone.call()

//     phone = new(IPhone)
//     phone.call()

// }

// package main

// import "fmt"

// type Shape interface {
//     area() float64
// }

// type Rectangle struct {
//     width  float64
//     height float64
// }

// func (r Rectangle) area() float64 {
//     return r.width * r.height
// }

// type Circle struct {
//     radius float64
// }

// func (c Circle) area() float64 {
//     return 3.14 * c.radius * c.radius
// }

// func main() {
//     var s Shape

//     s = Rectangle{width: 10, height: 5}
//     fmt.Printf("矩形面积: %f\n", s.area())

//     s = Circle{radius: 3}
//     fmt.Printf("圆形面积: %f\n", s.area())
// }

// package main

// import "fmt"

// func FindIndex[T comparable](slice []T, target T) int {
//     for i, v := range slice {
//         if v == target {
//             return i
//         }
//     }
//     return -1
// }
// func main(){
// // 使用示例
// numbers := []int{1, 2, 3, 4, 5}
// fmt.Println(FindIndex(numbers, 3))  // 输出: 2

// names := []string{"Alice", "Bob", "Charlie"}
// fmt.Println(FindIndex(names, "Bob"))  // 输出: 1
// }

// package main

// import "fmt"

// // 数字类型约束
// type Number interface {
//     int | int8 | int16 | int32 | int64 |
//     uint | uint8 | uint16 | uint32 | uint64 |
//     float32 | float64
// }

// func Add[T Number](a, b T) T {
//     return a + b
// }

// // 使用示例
// func main(){
// fmt.Println(Add(10, 20))        // 输出: 30
// fmt.Println(Add(3.14, 2.71))    // 输出: 5.85
// }

// package main

// import "fmt"

// // 定义 Stringer 约束
// type Stringer interface {
//     String() string
// }

// func PrintString[T Stringer](value T) {
//     fmt.Println(value.String())
// }

// // 实现自定义类型
// type Person struct {
//     Name string
//     Age  int
// }

// func (p Person) String() string {
//     return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
// }
// func main(){

// // 使用示例
// person := Person{Name: "Alice", Age: 25}
// PrintString(person)  // 输出: Alice (25 years old)
// }

// package main

// import "fmt"

// // 交换两个值
// func Swap[T any](a, b T) (T , T) {
//     return b, a
// }

// // 判断切片是否包含元素
// func Contains[T comparable](slice []T, target T) bool {
//     for _, item := range slice {
//         if item == target {
//             return true
//         }
//     }
//     return false
// }

// // 去重函数
// func Unique[T comparable](slice []T) []T {
//     seen := make(map[T]bool)
//     result := []T{}

//     for _, item := range slice {
//         if !seen[item] {
//             seen[item] = true
//             result = append(result, item)
//         }
//     }
//     return result
// }

// // 使用示例
// func main() {
//     // Swap 示例
//     a, b := 10, 20
//     a, b = Swap(a, b)
//     fmt.Printf("a=%d, b=%d\n", a, b)  // 输出: a=20, b=10

//     // Contains 示例
//     numbers := []int{1, 2, 3, 4, 5}
//     fmt.Println(Contains(numbers, 3))  // 输出: true

//     // Unique 示例
//     duplicates := []int{1, 2, 2, 3, 4, 4, 5}
//     unique := Unique(duplicates)
//     fmt.Println(unique)  // 输出: [1 2 3 4 5]
// }

// 求切片最大值
// func Max[T Number](slice []T) T {
//     if len(slice) == 0 {
//         var zero T
//         return zero
//     }

//     max := slice[0]
//     for _, value := range slice[1:] {
//         if value > max {
//             max = value
//         }
//     }
//     return max
// }

// // 求切片最小值
// func Min[T Number](slice []T) T {
//     if len(slice) == 0 {
//         var zero T
//         return zero
//     }

//     min := slice[0]
//     for _, value := range slice[1:] {
//         if value < min {
//             min = value
//         }
//     }
//     return min
// }

// // 求切片平均值
// func Average[T Number](slice []T) float64 {
//     if len(slice) == 0 {
//         return 0
//     }

//     var sum T
//     for _, value := range slice {
//         sum += value
//     }
//     return float64(sum) / float64(len(slice))
// }

// // 使用示例
// func main() {
//     ints := []int{1, 5, 3, 9, 2}
//     floats := []float64{1.1, 5.5, 3.3, 9.9, 2.2}

//     fmt.Printf("Max int: %d\n", Max(ints))           // 输出: 9
//     fmt.Printf("Min float: %.1f\n", Min(floats))     // 输出: 1.1
//     fmt.Printf("Average: %.2f\n", Average(floats))   // 输出: 4.40
// }

// package main

// import "fmt"

// // 泛型栈实现
// type Stack[T any] struct {
//     elements []T
// }

// // 入栈
// func (s *Stack[T]) Push(value T) {
//     s.elements = append(s.elements, value)
// }

// // 出栈
// func (s *Stack[T]) Pop() (T, bool) {
//     if len(s.elements) == 0 {
//         var zero T
//         return zero, false
//     }

//     lastIndex := len(s.elements) - 1
//     value := s.elements[lastIndex]
//     s.elements = s.elements[:lastIndex]
//     return value, true
// }

// // 查看栈顶元素
// func (s *Stack[T]) Peek() (T, bool) {
//     if len(s.elements) == 0 {
//         var zero T
//         return zero, false
//     }
//     return s.elements[len(s.elements)-1], true
// }

// // 判断栈是否为空
// func (s *Stack[T]) IsEmpty() bool {
//     return len(s.elements) == 0
// }

// // 使用示例
// func main() {
//     // 整数栈
//     intStack := Stack[int]{}
//     intStack.Push(1)
//     intStack.Push(2)
//     intStack.Push(3)

//     fmt.Println(intStack.Pop())  // 输出: 3 true

//     // 字符串栈
//     stringStack := Stack[string]{}
//     stringStack.Push("hello")
//     stringStack.Push("world")

//     fmt.Println(stringStack.Pop())  // 输出: world true
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // 线程安全的泛型映射
// type SafeMap[K comparable, V any] struct {
//     data map[K]V
//     mutex sync.RWMutex
// }

// // 创建新的 SafeMap
// func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
//     return &SafeMap[K, V]{
//         data: make(map[K]V),
//     }
// }

// // 设置键值对
// func (m *SafeMap[K, V]) Set(key K, value V) {
//     m.mutex.Lock()
//     defer m.mutex.Unlock()
//     m.data[key] = value
// }

// // 获取值
// func (m *SafeMap[K, V]) Get(key K) (V, bool) {
//     m.mutex.RLock()
//     defer m.mutex.RUnlock()
//     value, exists := m.data[key]
//     return value, exists
// }

// // 删除键
// func (m *SafeMap[K, V]) Delete(key K) {
//     m.mutex.Lock()
//     defer m.mutex.Unlock()
//     delete(m.data, key)
// }

// // 获取所有键
// func (m *SafeMap[K, V]) Keys() []K {
//     m.mutex.RLock()
//     defer m.mutex.RUnlock()

//     keys := make([]K, 0, len(m.data))
//     for key := range m.data {
//         keys = append(keys, key)
//     }
//     return keys
// }

// // 使用示例
// func main() {
//     // 创建字符串到整数的映射
//     scores := NewSafeMap[string, int]()
//     scores.Set("Alice", 95)
//     scores.Set("Bob", 87)

//     if score, exists := scores.Get("Alice"); exists {
//         fmt.Printf("Alice's score: %d\n", score)  // 输出: Alice's score: 95
//     }

//     fmt.Println("Keys:", scores.Keys())  // 输出: Keys: [Alice Bob]
// }

// package main

// import "fmt"

// // 你的实现代码在这里
// func Map[T any, U any](slice []T, mapper func(T) U) []U {
//     // 实现映射逻辑
// }

// // 测试代码
// func main() {
//     numbers := []int{1, 2, 3, 4, 5}
//     strings := Map(numbers, func(n int) string {
//         return fmt.Sprintf("Number: %d", n)
//     })
//     fmt.Println(strings)
// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {
//     err := errors.New("this is an error")
//     fmt.Println(err)
// }

// package main

// import (
// 		"fmt"
// )

// type DivideError struct {
//         Dividend int
//         Divisor  int
// }

// func (e *DivideError) Error() string {
//         return fmt.Sprintf("cannot divide %d by %d", e.Dividend, e.Divisor)
// }

// func divide(a, b int) (int, error) {
//         if b == 0 {
//                 return 0, &DivideError{Dividend: a, Divisor: b}
//         }
//         return a / b, nil
// }

// func main() {
//         _, err := divide(10, 0)
//         if err != nil {
//                 fmt.Println(err) // 输出：cannot divide 10 by 0
//         }
// }

// package main

// import (
// 	"fmt"
// )

// // 定义一个 DivideError 结构
// type DivideError struct {
//     dividee int
//     divider int
// }

// // 实现 `error` 接口
// func (de *DivideError) Error() string {
//     strFormat := `
//     Cannot proceed, the divider is zero.
//     dividee: %d
//     divider: 0
// `
//     return fmt.Sprintf(strFormat, de.dividee)
// }

// // 定义 `int` 类型除法运算的函数
// func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
//     if varDivider == 0 {
//             dData := DivideError{
//                     dividee: varDividee,
//                     divider: varDivider,
//             }
//             errorMsg = dData.Error()
//             return
//     } else {
//             return varDividee / varDivider, ""
//     }

// }

// func main() {

//     // 正常情况
//     if result, errorMsg := Divide(100, 10); errorMsg == "" {
//             fmt.Println("100/10 = ", result)
//     }
//     // 当除数为零的时候会返回错误信息
//     if _, errorMsg := Divide(100, 0); errorMsg != "" {
//             fmt.Println("errorMsg is: ", errorMsg)
//     }

// }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type MyError struct {
//         Code int
//         Msg  string
// }

// func (e *MyError) Error() string {
//         return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
// }

// func getError() error {
//         return &MyError{Code: 404, Msg: "Not Found"}
// }

// func main() {
//         err := getError()
//         var myErr *MyError
//         if errors.As(err, &myErr) {
//                 fmt.Printf("Custom error - Code: %d, Msg: %s\n", myErr.Code, myErr.Msg)
//         }
// }

// package main

// import "fmt"

// func safeFunction() {
//         defer func() {
//                 if r := recover(); r != nil {
//                         fmt.Println("Recovered from panic:", r)
//                 }
//         }()
//         panic("something went wrong")
// }

// func main() {
//         fmt.Println("Starting program...")
//         safeFunction()
//         fmt.Println("Program continued after panic")
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sayHello() {
//         for i := 0; i < 5; i++ {
//                 fmt.Println("Hello")
//                 time.Sleep(100 * time.Millisecond)
//         }
// }

// func main() {
//         go sayHello() // 启动 Goroutine
//         for i := 0; i < 5; i++ {
//                 fmt.Println("Main")
//                 time.Sleep(100 * time.Millisecond)
//         }
// }

// package main

// import "fmt"

// func sum(s []int, c chan int) {
//     sum := 0
//     for _, v := range s {
//         sum += v
//     }
//     c <- sum // 把 sum 发送到通道 c
// }

// func main() {
//     s := []int{7, 2, 8, -9, 4, 0}

//     c := make(chan int)
//     go sum(s[:len(s)/2], c)
//     go sum(s[len(s)/2:], c)
//     x, y := <-c, <-c // 从通道 c 中接收

//     fmt.Println(x, y, x+y)
// }

// package main

// import (
// 	"fmt"
// )

// func fibonacci(n int, c chan int) {
//     x, y := 0, 1
//     for i := 0; i < n; i++ {
//         c <- x
//         x, y = y, x+y
//     }
//     close(c)
// }

// func main() {
//     c := make(chan int, 10)
//     go fibonacci(cap(c), c)

//     for i := range c {
//         fmt.Println(i)
//     }
// }

// package main

// import "fmt"

// func fibonacci(c, quit chan int) {
//     x, y := 0, 1
//     for {
//         select {
//         case c <- x:
//             x, y = y, x+y
//         case <-quit:
//             fmt.Println("quit")
//             return
//         }
//     }
// }

// func main() {
//     c := make(chan int)
//     quit := make(chan int)

//     go func() {
//         for i := 0; i < 10; i++ {
//             fmt.Println(<-c)
//         }
//         quit <- 0
//     }()
//     fibonacci(c, quit)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, wg *sync.WaitGroup) {
//         defer wg.Done()
//         fmt.Printf("Worker %d started\n", id)
//         fmt.Printf("Worker %d finished\n", id)
// 	}

// func main() {
//         var wg sync.WaitGroup

//         for i := 1; i <= 3; i++ {
//                 wg.Add(1)
//                 go worker(i, &wg)
//         }

//         wg.Wait()
//         fmt.Println("All workers done")
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
//         // 创建文件，如果文件已存在会被截断（清空）
//         file, err := os.Create("test.txt")
//         if err != nil {
//                 log.Fatal(err)
//         }
//         defer file.Close() // 确保文件关闭

//         log.Println("文件创建成功")
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
//     file, err := os.Create("test.txt")
//     if err != nil {
//         fmt.Println("Error creating file:", err)
//         return
//     }
//     defer file.Close()

//     writer := bufio.NewWriter(file)
//     fmt.Fprintln(writer, "Hello, World!")
//     writer.Flush()
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
//     file, err := os.Create("test.txt")
//     if err != nil {
//         fmt.Println("Error creating file:", err)
//         return
//     }
//     defer file.Close()

//     writer := bufio.NewWriter(file)
//     fmt.Fprintln(writer, "Hello!")
//     writer.Flush()
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
//     file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
//     if err != nil {
//         fmt.Println("Error opening file:", err)
//         return
//     }
//     defer file.Close()

//     if _, err := file.WriteString("Appended text\n"); err != nil {
//         fmt.Println("Error appending to file:", err)
//         return
//     }

//     fmt.Println("Text appended successfully!")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
//     err := os.Remove("test.txt")
//     if err != nil {
//         fmt.Println("Error deleting file:", err)
//         return
//     }

//     fmt.Println("File deleted successfully!")
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
//         if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
//                 fmt.Println("文件不存在")
//         } else {
//                 fmt.Println("文件存在")
//         }
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
//         // 创建单个目录
//         err := os.Mkdir("newdir", 0755)
//         if err != nil {
//                 log.Fatal(err)
//         }

//         // 递归创建多级目录
//         // err = os.MkdirAll("path/to/newdir", 0755)
//         // if err != nil {
//         //         log.Fatal(err)
//         // }
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
//         entries, err := os.ReadDir(".")
//         if err != nil {
//                 log.Fatal(err)
//         }

//         for _, entry := range entries {
//                 info, _ := entry.Info()
//                 fmt.Printf("%-20s %8d %v\n",
//                         entry.Name(),
//                         info.Size(),
//                         info.ModTime().Format("2006-01-02 15:04:05"))
//         }
// }

// package main

// import (
// 	"log"
// 	"os"
// )

// func main() {
//         // 删除空目录
//         err := os.Remove("newdir")
//         if err != nil {
//                 log.Fatal(err)
//         }

//         // 递归删除目录及其内容
//         // err = os.RemoveAll("path/to/dir")
//         // if err != nil {
//         //         log.Fatal(err)
//         // }
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
//         // 创建临时文件
//         tmpFile, err := os.CreateTemp("", "example-*.txt")
//         if err != nil {
//                 log.Fatal(err)
//         }
//         defer os.Remove(tmpFile.Name()) // 清理

//         fmt.Println("临时文件:", tmpFile.Name())

//         // 创建临时目录
//         tmpDir, err := os.MkdirTemp("", "example-*")
//         if err != nil {
//                 log.Fatal(err)
//         }
//         defer os.RemoveAll(tmpDir) // 清理

//         fmt.Println("临时目录:", tmpDir)
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
//     pattern := `^[a-zA-Z0-9]+$`
//     regex := regexp.MustCompile(pattern)

//     str := "Hello123"
//     if regex.MatchString(str) {
//         fmt.Println("字符串匹配正则表达式")
//     } else {
//         fmt.Println("字符串不匹配正则表达式")
//     }
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
//     pattern := `\d+`
//     regex := regexp.MustCompile(pattern)

//     str := "我有 3 个苹果和 5 个香蕉"
//     matches := regex.FindAllString(str, -1)
//     fmt.Println("找到的数字：", matches)
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
//     pattern := `,`
//     regex := regexp.MustCompile(pattern)

//     str := "apple,banana,orange"
//     parts := regex.Split(str, -1)
//     fmt.Println("分割后的字符串：", parts)
// }

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
//     pattern := `\s+`
//     regex := regexp.MustCompile(pattern)

//     str := "Hello    World"
//     result := regex.ReplaceAllString(str, " ")
//     fmt.Println("替换后的字符串：", result)
// }

// package main

// import "fmt"

// func main() {
//     var i interface{} = "Hello, Go!"

//     // 尝试将 i 断言为 string 类型
//     s, ok := i.(string)
//     if ok {
//         fmt.Println("断言成功:", s)
//     } else {
//         fmt.Println("断言失败")
//     }

//     // 尝试将 i 断言为 int 类型
//     n, ok := i.(int)
//     if ok {
//         fmt.Println("断言成功:", n)
//     } else {
//         fmt.Println("断言失败")
//     }
// }

// package main

// import "fmt"

// func main() {
//     var i interface{} = "Hello, Go!"

//     // 直接断言为 string 类型
//     s := i.(string)
//     fmt.Println("断言成功:", s)

//     // 直接断言为 int 类型（会引发 panic）
//     n := i.(int)
//     fmt.Println("断言成功:", n)
// }

// package main

// import "fmt"

// // 父结构体
// type Animal struct {
//     Name string
// }

// // 父结构体的方法
// func (a *Animal) Speak() {
//     fmt.Println(a.Name, "says hello!")
// }

// // 子结构体
// type Dog struct {
//     Animal // 嵌入 Animal 结构体
//     Breed  string
// }

// func main() {
//     dog := Dog{
//         Animal: Animal{Name: "Buddy"},
//         Breed:  "Golden Retriever",
//     }

//     dog.Speak() // 调用父结构体的方法
//     fmt.Println("Breed:", dog.Breed)
// }

// package main

// import "fmt"

// // 定义接口
// type Speaker interface {
//     Speak()
// }

// // 父结构体
// type Animal struct {
//     Name string
// }

// // 实现接口方法
// func (a *Animal) Speak() {
//     fmt.Println(a.Name, "says hello!")
// }

// // 子结构体
// type Dog struct {
//     Animal
//     Breed string
// }

// func main() {
//     var speaker Speaker

//     dog := Dog{
//         Animal: Animal{Name: "Buddy"},
//         Breed:  "Golden Retriever",
//     }

//     speaker = &dog
//     speaker.Speak() // 通过接口调用方法
// }

package main

import "fmt"

// 基类
type Vehicle struct {
    Brand string
}

func (v *Vehicle) Start() {
    fmt.Println(v.Brand, "started")
}

// 派生类
type Car struct {
    Vehicle // 嵌入Vehicle
    Model  string
}

// 重写Start方法
func (c *Car) Start() {
    fmt.Println(c.Brand, c.Model, "car started")
}

func main() {
    v := Vehicle{Brand: "Toyota"}
    c := Car{
        Vehicle: Vehicle{Brand: "Honda"},
        Model:  "Civic",
    }
    
    v.Start() // Toyota started
    c.Start() // Honda Civic car started
    c.Vehicle.Start() // Honda started
}