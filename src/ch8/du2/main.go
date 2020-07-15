package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0{
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func(){
		for _, root := range roots{
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the result periodically
	tick := time.Tick(500 * time.Millisecond)

	// print the results
	var nfiles, nbytes int64

	loop:
		for{
			select {
			case size, ok := <- fileSizes:
				if !ok{
					break loop
				}
				nfiles++
				nbytes += size
			case <- tick:
				printDiskUsage(nfiles, nbytes)

			}
		}

	printDiskUsage(nfiles, nbytes)


}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(dir string, in chan <- int64){
	for _, entry := range dirents(dir){
		if entry.IsDir(){
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, in)
		}else{
			in <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo{
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stdout, "du: %v\n", err)
		return nil
	}
	return entries
}