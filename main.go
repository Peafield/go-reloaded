package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getTextSlice() []string {
	body, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(string(body), " ")
}

func writeToFile(s string) {
	err := os.WriteFile("result.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func selector(e []string) []string {
	for i, ch := range e {
		if ch[:1] == "(" {
			switch {
			case ch == "(hex)":
				e[i-1] = hexAndBin(e[i-1], 16)
				e = remove(e, i)
			case ch == "(bin)":
				e[i-1] = hexAndBin(e[i-1], 2)
				e = remove(e, i)
			case ch[:2] == "(u":
				if ch[3:] == "," {
					a, err := strconv.Atoi(e[i+1][:1])
					if err != nil {
						fmt.Println("Error:", err)
					}
					for a > 0 {
						e[i-a] = up(e[i-a])
						a--
					}
					sum := 2
					for sum > 0 {
						e = remove(e, i)
						sum--
					}
				} else {
					e[i-1] = cap(e[i-1])
					e = remove(e, i)
				}
			case ch[:2] == "(l":
				if ch[4:] == "," {
					a, err := strconv.Atoi(e[i+1][:1])
					if err != nil {
						fmt.Println("Error:", err)
					}
					for a > 0 {
						e[i-a] = low(e[i-a])
						a--
					}
					sum := 2
					for sum > 0 {
						e = remove(e, i)
						sum--
					}

				} else {
					e[i-1] = cap(e[i-1])
					e = remove(e, i)
				}
			case ch[:2] == "(c":
				if ch[4:] == "," {
					a, err := strconv.Atoi(e[i+1][:1])
					if err != nil {
						fmt.Println("Error:", err)
					}
					for a > 0 {
						e[i-a] = cap(e[i-a])
						a--
					}
					sum := 2
					for sum > 0 {
						e = remove(e, i)
						sum--
					}

				} else {
					e[i-1] = cap(e[i-1])
					e = remove(e, i)
				}
			}
		}
		switch {
		case ch == "a":
			if isVowel(e[i+1][:1]) {
				e[i] = "an"
			}
		case ch == "an":
			if !(isVowel(e[i+1][:1])) {
				e[i] = "a"
			}
		}
	}
	return e
}

func hexAndBin(s string, i int) string {
	hexText, err := strconv.ParseInt(s, i, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return strconv.Itoa(int(hexText))
}

func up(s string) string {
	return strings.TrimSpace(strings.ToUpper(s))
}

func low(s string) string {
	return strings.ToLower(s)
}

func cap(s string) string {
	return strings.Title(s)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func converToString(s []string) string {
	return strings.Join(s, " ")
}

func isVowel(s string) bool {
	b := []rune(s)
	result := false
	for _, ch := range b {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			result = true
		}
	}
	return result
}

func main() {
	e := getTextSlice()
	writeToFile(converToString(selector(e)))
}

//TO DO: resolve issue of extra space after e.g. "...word (cap, 6) ,..." - results in "...word  , ..."
