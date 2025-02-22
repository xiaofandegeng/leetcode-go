## go面试题总结

### 1. 第一题
输出下面代码的结果

```go
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	panic("触发异常")
}
```

看下答案，输出：

```go
打印后
打印中
打印前
panic: 触发异常

```

参考解析：defer 的执行顺序是后进先出。当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic

### 2. 第二题
下面这段代码输出什么，说明原因。
```go
func main() {

    slice := []int{0, 1, 2, 3}
    m := make(map[int]*int)
    
    for key, val := range slice {
        m[key] = &val
    }
    
    for k, v := range m {
    fmt.Println(k, "->", *v)
    }
}

```
直接给答案：
```go
    0 -> 3
    1 -> 3
    2 -> 3
    3 -> 3

```
参考解析：这是新手常会犯的错误写法，for range 循环的时候会创建每个元素的副本，而不是元素的引用，所以 m[key] = &val 取的都是变量 val 的地址，所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3.

正确的写法:
```go
func main() {

     slice := []int{0,1,2,3}
     m := make(map[int]*int)

     for key,val := range slice {
         value := val
         m[key] = &value
     }

    for k,v := range m {
        fmt.Println(k,"===>",*v)
    }
}
```