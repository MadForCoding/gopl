package main

import(
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1: %v\n", err)
		os.Exit(1)
	}
	for _, links := range visit(nil, doc){
		fmt.Println(links)
	}
}


// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string{
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a"{
		for _, a := range n.Attr{
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}


	// iterate
	//for c := n.FirstChild; c != nil; c = c.NextSibling{
	//	links = visit(links, c)
	//}

	// recursive  exercise5-1
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}