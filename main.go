package main

import (
	"encoding/json"
	"fmt"
	"github.com/simultechnology/myproglog/internal/server"
	"log"
	"strings"
)

func main() {
	fmt.Println("start!")

	myLog := server.NewLog()

	jsonStr := `{"record": {"value": "ishi"}}`
	// strings.NewReaderを使用してio.Readerを作成
	reader := strings.NewReader(jsonStr)

	var req server.ProduceRequest
	err := json.NewDecoder(reader).Decode(&req)
	if err != nil {
		panic(err) // エラーハンドリングは実際の要件に応じて適切に行ってください
	}

	_, err = myLog.Append(req.Record)
	if err != nil {
		return
	}
	message, err := myLog.AppendMessage("simul")
	if err != nil {
		return
	}
	fmt.Println(message)

	text := "Now we know AVL trees offer O(myLog n) search performance"
	words := strings.Split(text, " ")
	for _, word := range words {
		myLog.AppendMessage(word)
	}

	fmt.Printf("%v\n", myLog)

	content, err := myLog.Read(5)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%v", string(content.Value))

	srv := server.NewHTTPServer(":58888", myLog)
	log.Fatal(srv.ListenAndServe())
}
