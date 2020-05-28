package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numberSign("-1242.123"))
}

func comma(number string) string{
	if len(number) <= 3 {
		return number
	}

	var buf bytes.Buffer
	headNum := len(number) % 3
	// Deal with head of number
	if headNum != 0 {
		for i := 0; i < headNum; i ++ {
			fmt.Fprintf(&buf, "%c", number[i])
		}
		buf.WriteByte(',')
	}

	// Rest of the number
	for i := headNum; i < len(number); i++ {
		fmt.Fprintf(&buf, "%c", number[i])
		if (i+1-headNum) % 3 == 0 && i != len(number) - 1{
			buf.WriteByte(',')
		}
	}

	return buf.String()
}



// commaFloat handles the float number and int number
func commaFloat(number string) string{
	if len(number) <= 3 {
		return number
	}
	if pos := strings.Index(number, "."); pos == -1 {
		return comma(number)
	} else if pos == 0{
		// case 0.12345
		return number
	}

	pos := strings.Index(number, ".")
	numberStr := number[:pos]
	decimalDot := number[pos:]
	var buf bytes.Buffer
	buf.WriteString(comma(numberStr))
	buf.WriteString(decimalDot)
	return buf.String()
}


func numberSign(number string) string {
	var buf bytes.Buffer
	if number == "" {
		return number
	}

	if number[0] == '+' {
		return commaFloat(number[1:])
	}else if number[0] == '-' {
		buf.WriteByte('-')
		buf.WriteString(commaFloat(number[1:]))
		return buf.String()
	}else{
		return commaFloat(number)
	}
}
