package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)

	key := flag.String("k", "", "key to decode or encode with")
	enc := flag.Bool("e", false, "encode instead of decode input")
	file := flag.String("f", "", "file to decode or encode (if not specified, stdin is assumed)")
	flag.Parse()

	input := ""
	if file != nil && len(*file) > 0 {
		data, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Fatal(err)
		}
		input = string(data)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input += scanner.Text()
		}
	}

	if len(input) == 0 {
		log.Fatal("no data to encode or decode provided")
	}

	if key == nil || len(*key) == 0 {
		log.Fatal("no key provided")
	}

	k := cleanKey(*key)

	dir, result, i, j := -1, []byte{}, 0, 0
	if *enc {
		dir = 1
	}

	for i < len(input) {
		n := input[i]
		if isLower(n) || isUpper(n) {
			keyVal := k[j] * dir
			newVal := n + byte(keyVal)
			if isUpper(n) && newVal > 'Z' {
				newVal = 'A' + newVal - 'Z' - byte(1)
			} else if isLower(n) && newVal > 'z' {
				newVal = 'a' + newVal - 'z' - byte(1)
			} else if isUpper(n) && newVal < 'A' {
				newVal = ('A' - newVal) + 'A'
			} else if isLower(n) && newVal < 'a' {
				newVal = ('a' - newVal) + 'a'
			}
			result = append(result, newVal)
			j++
			if j >= len(k) {
				j = 0
			}
		} else {
			result = append(result, n)
		}
		i++
	}

	os.Stdout.Write(result)
}

func cleanKey(provided string) []int {
	clean := []int{}
	for i := 0; i < len(provided); i++ {
		c := provided[i]
		if isLower(c) {
			clean = append(clean, int(c-'a'))
		} else if isUpper(c) {
			clean = append(clean, int(c-'A'))
		}
	}
	return clean
}

func isLower(n byte) bool {
	return n >= 'a' && n <= 'z'
}

func isUpper(n byte) bool {
	return n >= 'A' && n <= 'Z'
}
