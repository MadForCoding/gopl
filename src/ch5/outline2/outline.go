package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil{
		return err
	}

	forEachNode(doc, startElement, endElement)

	return nil

}

func forEachNode(n *html.Node, pre func(n *html.Node) bool, post func(n *html.Node)){
	var isEnd bool
	if pre != nil {
		isEnd = pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling{
		forEachNode(c, pre, post)
	}

	if post != nil && !isEnd{
		post(n)
	}

}

var depth int
// %*s has two arguments, first argument is the number of indent, the second argument is the prefix
func startElement(n *html.Node) bool{
	var isEnd bool
	if n.Type == html.ElementNode {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%*s<%s", depth*2, "", n.Data))
		if len(n.Attr) == 0{
			buf.WriteString("/")
			isEnd = true
		}else{
			for _,att := range n.Attr{
				buf.WriteString(fmt.Sprintf(" %s=%#v", att.Key,att.Val))
			}
			depth++
		}
		buf.WriteString(">")
		fmt.Println(buf.String())

	} else if n.Type == html.CommentNode{
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("%*s//%s", depth*2,"",n.Data))
		fmt.Println(buf.String())
		//depth++
	}

	return isEnd
}

func endElement(n *html.Node){
	if n.Type == html.ElementNode{
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		//depth--
	}
}

