package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	//readWordsFromStdin()
	//readWordsFromText()
	var words WordCounter
	fmt.Fprintf(&words, "append something to the end")
	fmt.Println(words)
	words = 0
	words.Write([]byte("Hello this is a text for testing how many words you can get"))
	fmt.Println(words)
}

type WordCounter int

// count how many words
func (w *WordCounter) Write(buf []byte) (int, error){
	for start := 0; start < len(buf); {
		advance, _, err := bufio.ScanWords(buf[start:], true)
		if err != nil{
			log.Fatal(err)
		}
		start += advance
		*w++
	}
	return int(*w), nil
}


// readWordsFromText count the frequency of word by the function of bufio.ScanWords
func readWordsFromText() map[string]int{
	plainText := "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed " +
		"ullamco, leberkas sint short loin strip steak ut shoulder " +
		"shankle porchetta venison prosciutto turducken swine.\n " +
		"Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"

	counts := make(map[string]int)
	input := bufio.NewScanner(strings.NewReader(plainText))
	input.Split(bufio.ScanWords)

	for input.Scan(){
		counts[input.Text()]++
	}

	for key, value := range counts{
		fmt.Printf("%s : %d\n", key, value)
	}

	return counts
}


// tradional function to count the frequency of word from Stdin
// use Scanner default split function ScanLines
func readWordsFromStdin() map[string]int{
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan(){
		if input.Text() == "end"{
			break
		}
		counts[input.Text()]++
	}

	for key, value := range counts{
		if value > 1{
			fmt.Printf("Word: %s appears %d times\n", key, value)
		}
	}
	fmt.Println(counts)

	return counts

}
