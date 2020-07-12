package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct{
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main(){
	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	// Reverse
	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nReverse(byYear)")
	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)

	// Custom sort
	fmt.Println("\nCustom sort")
	sort.Sort(CustomSort{tracks, func(x, y *Track) bool{
		if x.Title != y.Title{
			return x.Title < y.Title
		}
		if x.Year != y.Year{
			return x.Year < y.Year
		}
		if x.Length != y.Length{
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)


	// check the slice is sorted
	fmt.Println("\nBelow is the function check the slice is sorted")
	var s = []int{3,1,4,2}
	fmt.Println(sort.IntsAreSorted(s))
	fmt.Println("sort the []int")
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s))
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	fmt.Println(s)
	fmt.Println(sort.IntsAreSorted(s))
}


func length(s string) time.Duration{
	d, err := time.ParseDuration(s)
	if err != nil{
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}



// sort by artist
type byArtist []*Track

func (b byArtist) Len() int{
	return len(b)
}

func (b byArtist) Less(i, j int) bool{
	return b[i].Artist < b[j].Artist
}

func (b byArtist) Swap(i, j int){
	b[i], b[j] = b[j], b[i]
}

// sort by year
type byYear []*Track

func (b byYear) Len() int{
	return len(b)
}

func (b byYear) Less(i, j int) bool{
	return b[i].Year < b[j].Year
}

func (b byYear) Swap(i, j int){
	b[i], b[j] = b[j], b[i]
}


// Custom struct sorting by the title, year, length
type CustomSort struct{
	t []*Track
	less func(x, y *Track) bool
}

func (c CustomSort) Len() int{
	return len(c.t)
}

func (c CustomSort) Less(i, j int) bool{
	return c.less(c.t[i], c.t[j])
}

func (c CustomSort) Swap(i, j int){
	c.t[i], c.t[j] = c.t[j], c.t[i]
}





