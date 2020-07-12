package main

import (
	"ch7/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlagFunc("temp", 20.0, "the temperature")

func main(){

	// flag.Parse()会根据命令行传递过来的参数
	// 查找comandline中的变量是否和参数相对应，若对应，会调用变量中的Set方法进行传递
	flag.Parse()
	fmt.Println(*temp)



}
