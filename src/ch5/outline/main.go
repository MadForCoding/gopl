package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

var hashmap = make(map[string]int)

func main(){
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}

	outline(nil, doc)
	fmt.Println("--------------")
	for k,v := range hashmap{
		fmt.Println(k,": ",v)
	}
}

func outline(stack []string, n *html.Node){
	if n.Type == html.ElementNode {

		stack = append(stack, n.Data)
		record(n.Data)
		fmt.Println(stack)
		if n.Data != "script" && n.Data != "style" {
			printContent(stack, n.Attr)
		}

	}

	for c := n.FirstChild; c !=nil ; c = c.NextSibling{

		outline(stack, c)
	}
}

// exercise 5-2
func record(data string){
	hashmap[data]++
}

// exercise 5-3
func printContent(stack []string, attr []html.Attribute) {
	for _, elem:= range attr{
		fmt.Printf("%v: attr: %s, value: %s\n", stack, elem.Key,elem.Val)
	}
}
