package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type editor struct {
	text []string
}

func getTextSlice() []string {
	body, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(string(body), " ")
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// TO DO: Range through []string outside of any particular function.
func hexAndBin(e []string) []string {
	for i, ch := range e {
		if ch == "(hex)" {
			hexText, err := strconv.ParseInt(e[i-1], 16, 64)
			if err != nil {
				log.Fatalf("Error %v", err)
			}
			e[i-1] = strconv.Itoa(int(hexText))
			remove(e, i)
		}
		if ch == "(bin)" {
			hexText, err := strconv.ParseInt(e[i-1], 2, 64)
			if err != nil {
				log.Fatalf("Error %v", err)
			}
			e[i-1] = strconv.Itoa(int(hexText))
			remove(e, i)
		}
	}
	return e
}

func main() {
	e := editor{}
	e.text = getTextSlice()
	fmt.Println(hexAndBin(e.text))
}
