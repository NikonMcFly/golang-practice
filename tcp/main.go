package main


import (
	"fmt"
	"net"
	"bufio"
	"log"
	"io"
	"strings"
)



type Command struct {
	Fields []string
	Result chan string
}

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil{
		log.Fatalln("err: ", err)
	}
	defer conn.Close()

	commands := make(chan Command)
	go redisServer(commands)

	for {
		ln, err := conn.Accept()
		if err != nil{
			fmt.Println("conn broke: ", err)
		}
		go handle(commands, ln)
	}
}

func handle(commands chan Command, conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		result := make(chan string)
		commands <- Command{
			Fields: fs,
			Result: result,
		}

		io.WriteString(conn, <-result+"\n")
	}
}

func redisServer(commands chan Command) {
	var data = make(map[string]string)
	for cmd := range commands {
		if len(cmd.Fields) < 2 {
			cmd.Result <- "Expected at least 2 arguments"
			continue
		}

		fmt.Println("GOT Command", cmd)

	switch cmd.Fields[0] {
		case "GET":
			key := cmd.Fields[1]
			value := data[key]

			cmd.Result <- value
		
	case "SET":
		if len(cmd.Fields) != 3 {
			cmd.Result <- "EXPECTED VALUE"
			continue
		}

		key := cmd.Fields[1]
		value := cmd.Fields[2]
		data[key] = value
		cmd.Result <- ""

	case "DEL":
		key := cmd.Fields[1]
		delete(data, key)
		cmd.Result <- ""
	default:
		cmd.Result <- "INVALID COMMAND " + cmd.Fields[0] + "\n"
		}
	}
}