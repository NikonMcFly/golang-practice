package main

import (
	// "io"
	"fmt"
	"io/ioutil"
	"strings"
	"bytes"
)

func main() {

	z := []byte("steve again")
	k := bytes.NewReader(z)
	b, _ := ioutil.ReadAll(k)
	fmt.Println(string(b), k)

	reader := strings.NewReader("source string yo")

	var first_word string
	var second_word string
	var third_word string

	n, err := fmt.Fscan(reader, &first_word, &second_word, &third_word)
	fmt.Println(err, n)

	fmt.Println("first word", first_word)
	fmt.Println("second word", second_word)
	fmt.Println("third word", third_word)
}
