package tempconv

import (
	"ch2/tempconv"
	"flag"
	"fmt"
)

// CelsiusFlag 实现了 Value接口（String() 和 Set())
type CelsiusFlag struct{
	tempconv.Celsius
}

// 该方法中的实参由命令行参数传递过来
// 然后设置结构体中的值
func (f *CelsiusFlag) Set(s string) error{
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit{
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "K":
		f.Celsius = tempconv.KtoC(tempconv.Kelvin(value))
	}

	return fmt.Errorf("invalid temperature %q", s)
}


func CelsiusFlagFunc(name string, value tempconv.Celsius, usage string) *tempconv.Celsius{
	// 此时，结构体中的值为默认值20
	f := CelsiusFlag{value}
	//fmt.Println(f.Celsius)
	// 将参数为name的添加到命令行参数中
	// 随后flag.parse的时候，会对结构体中的值根据命令行传递过来的值做修改
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}


