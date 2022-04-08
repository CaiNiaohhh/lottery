package other

// 只是声明变量后将指针加到slice里面并不会发生内存逃逸，
// 但是如果有用append,fmt.Println,make这些的话就会发生内存逃逸
