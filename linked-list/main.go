package main


import (
    "fmt"
    "strings"
    "os"
    "strconv"
    "io/ioutil"
)

type student struct {
	name string
	age int
	num string
	next *student
}


func main() {
	students := new(student)
	
	students.next = nil

	if len(os.Args) < 2 {
		fmt.Println("no file to use")
		return
	}

	filename := os.Args[1]

	data, _ := ioutil.ReadFile(filename)

	for _, line := range strings.Split(string(data), "\n") {

		if line == "" {
			continue
		}

		td := strings.Split(string(line), " ")

		ts := &student {
			name : td[0],
			num : td[2],
			next : students.next,
		}

		ts.age, _ = strconv.Atoi(td[1])

		students.next = ts
	}

	for s := students.next; s != nil; s = s.next {
		fmt.Println(s.name, s.age, s.num)
	}
}

