package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
	out, count := CounterWriter(os.Stdout)
	fmt.Fprintf(out, "Print something to the screen\n")
	fmt.Println(*count)

}

type WordsCount struct{
	Writer io.Writer
	Count int
}

func (w *WordsCount) Write(buf []byte) (int, error){
	n, err := w.Writer.Write(buf)
	if err != nil{
		return n, err
	}
	w.Count += n
	return n, nil
}

func CounterWriter(writer io.Writer) (io.Writer, *int){
	w := WordsCount{
		Writer: writer,
	}
	return &w, &w.Count
}
