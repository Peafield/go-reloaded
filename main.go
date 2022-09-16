package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// getTextSlice creates a string slice of the text from 'sample.txt'
func getTextSlice() []string {
	body, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return strings.Split(string(body), " ")
}

// writeToFile writes the completed string to a file called 'result.txt'
func writeToFile(s string) {
	err := os.WriteFile("result.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// selector loops through a string slice and applies different functions depending on input.
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
		case ch == "A":
			if isVowel(e[i+1][:1]) {
				e[i] = "An"
			}
		case ch == "an":
			if !(isVowel(e[i+1][:1])) {
				e[i] = "a"
			}
		case ch == "An":
			if !(isVowel(e[i+1][:1])) {
				e[i] = "A"
			}
		}
		if ch == "," {
			fmt.Println("Found")
		}
	}
	return e
}

// hexAndBin converts a string to it's hex or bin value and returns it as a string
func hexAndBin(s string, i int) string {
	hexText, err := strconv.ParseInt(s, i, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return strconv.Itoa(int(hexText))
}

// up makes a string uppercase
func up(s string) string {
	return strings.ToUpper(s)
}

// low makes a string lowercase
func low(s string) string {
	return strings.ToLower(s)
}

// cap capitalise a string
func cap(s string) string {
	return strings.Title(s)
}

// isVowel checks with a string is a vowel or not
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

func isPunctuation(c rune) bool {
	delim := false
	if c >= 33 && c <= 47 {
		delim = true
	}
	return delim
}

// remove removes an element from a slice
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func insert(r []rune, i int) []rune {
	r = append(r[:i+1], r[i:]...)
	r[i] = 32
	return r
}

// convertToString converts a string slice into a string
func converToString(s []string) string {
	return strings.Join(s, " ")
}

// TO DO: Add a space after punctuation if necessary e.g. "folder,do...""
// converToRuneRemoveSpace converts a string to a slice of rune and then deletes white space
func converToRuneRemoveOrAddSpace(s string) []rune {
	r := []rune(s)
	for i, ch := range r {
		// Deleting white space before all puncutation.
		if (i < len(r)-1 && isPunctuation(ch) && r[i-1] == 32) || (i == len(r)-1 && isPunctuation(ch) && r[i-1] == 32) {
			r = append(r[:i-1], r[i:]...)
		}
		// Deleting white space after apostrophes
		if i < len(r)-1 && ch == 39 && r[i+1] == 32 {
			r = append(r[:i+1], r[i+2:]...)
		}
		// Inserting whitespace after commas
		if i < len(r)-1 && ch == 44 && r[i+1] != 32 {
			r = insert(r, i)
		}

	}
	return r
}

func main() {
	e := getTextSlice()
	s := selector(e)
	str := converToString(s)
	writeToFile(string(converToRuneRemoveOrAddSpace(str)))
}
