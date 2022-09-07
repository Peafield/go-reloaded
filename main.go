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

func selector(e []string) []string {
	for i, ch := range e {
		switch {
		case ch == "(hex)":
			e[i-1] = hexAndBin(e[i-1], 16)
			remove(e, i)
		case ch == "(bin)":
			e[i-1] = hexAndBin(e[i-1], 2)
			remove(e, i)
		}
	}
	return e
}

func hexAndBin(s string, i int) string {
	hexText, err := strconv.ParseInt(s, i, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	hexed := strconv.Itoa(int(hexText))
	return hexed
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	e := editor{}
	e.text = getTextSlice()
	fmt.Println(selector(e.text))
}
