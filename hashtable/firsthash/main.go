package main

import (
	"fmt"
	"log"
	"net/http"
	//"io/ioutil"
	"bufio"
	"os"
)

func main() {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""

	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading Input:", err)

	}
	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++
	}
}
