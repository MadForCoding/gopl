package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main(){
	// 类型断言 x.(T)
	// 第一种情况, 如果T为具体的类型，类型断言会检测x的动态类型是否和T相同
	// 若相同， 返回x的动态值。
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	//c := w.(*bytes.Buffer)
	fmt.Printf("%T\n", f)
	//fmt.Println(c)

	var wr io.ReadWriter
	wr = new(bytes.Buffer)
	check := wr.(*bytes.Buffer)
	fmt.Printf("%T\n", check)

	// 第二种情况，如果相反地断言的类型T是一个接口类型，然后类型断言检查是否x的动态类型满足T。
	//如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同动态类型和值部分的接口值，但是结果为类型T
	var w2 io.Writer
	w2 = os.Stdout
	rw := w2.(io.ReadWriter) // success: *os.File has both Read and Write
	//w2 = new(ByteCounter)
	rw = w2.(io.ReadWriter) // panic: *ByteCounter has no Read method
	w2 = rw.(io.Writer)
	fmt.Println(w2, rw)

}
