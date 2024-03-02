package main

import (
	"fmt"
	"github.com/simultechnology/myproglog/internal/server"
	"strings"
)

func main() {
	fmt.Println("start!")

	log := server.NewLog()

	record := server.Record{
		Value: []byte("log..."),
	}
	_, err := log.Append(record)
	if err != nil {
		return
	}
	message, err := log.AppendMessage("simul")
	if err != nil {
		return
	}
	fmt.Println(message)

	text := "Now we know AVL trees offer O(log n) search performance"
	words := strings.Split(text, " ")
	for _, word := range words {
		log.AppendMessage(word)
	}

	fmt.Printf("%v\n", log)

	content, err := log.Read(5)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%v", string(content.Value))
}
