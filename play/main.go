package main

import (
		"fmt"
		// "strconv"
		"os"
)



func main(){
	b, err := os.Create("text.html")
	// k, err := os.Open("text.html")
		if err != nil {
			fmt.Println("something with wrong opening")
		}
		// k, _ := os.Open("text.html")
		s := "<h1>steve is winning</h1>"
		g, _ := b.WriteString(s)
		fmt.Println(g)

		z, _ := os.Stat("text.html")
		fmt.Println(z.Size())
		os.Clearenv()
}
