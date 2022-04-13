package main

import "fmt"

// 只是声明变量后将指针加到slice里面并不会发生内存逃逸，
// 但是如果有用append,fmt.Println,make这些的话就会发生内存逃逸


// 闭包，内存逃逸
func Foo () func () int {
	x := 5            // x发生逃逸，因为在Foo调用完成后，被闭包函数用到，还不能回收，只能放到堆上存放
	return func () int {
		x += 1
		return x
	}
}

// Slice指针 []*int
func A() {
	var x int
	x = 10
	var ls []*int
	ls = append(ls, &x)        // x发生逃逸，ls存储的是指针，所以ls底层的数组虽然在栈存储，但x本身却是逃逸到堆上
}

// 切片扩容后长度太大，导致栈空间不足，逃逸到堆上。
func B() {
	s := make([]int, 10000, 10000)
	for index, _ := range s {
		s[index] = index
	}
}

//在 interface 类型上调用方法。 在 interface 类型上调用方法时会把interface变量使用堆分配， 因为方法的真正实现只能在运行时知道。
type foo interface {
	fooFunc()
}
type foo1 struct{}
func (f1 foo1) fooFunc() {}
func C() {
	var f foo
	f = foo1{}
	f.fooFunc()   // 调用方法时，f发生逃逸，因为方法是动态分配的
}

func main() {
	inner := Foo()
	fmt.Println(inner())
	fmt.Println(inner())
}