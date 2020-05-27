// Fetch prints the content found at the URL
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// practice1-8
		if !strings.HasPrefix(url, "http://") {
			var buf bytes.Buffer
			buf.WriteString("http://")
			buf.WriteString(url)
			url = buf.String()
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Oringinal method
		//b, err := ioutil.ReadAll(resp.Body)

		// practice1-7
		_, err = io.Copy(os.Stdout, resp.Body)


		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// practice1-9
		fmt.Printf("%d", resp.StatusCode)
	}
}
