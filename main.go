package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	notificationTitle := os.Args[1]
	if len(notificationTitle) == 0{
		err := fmt.Errorf("Please enter the intofictional title .")
		log.Fatal(err)
	}
	url := "https://server.notes.wuchuheng.com/graphql"
	var jsonStr = []byte(fmt.Sprintf(`
		{
		  "operationName": "CREATE_TODO",
		  "variables": {
			"input": {
			  "title": "%s"
			}
		  },
		  "query": "mutation CREATE_TODO($input: InputCreateTodoInput!) {\n  createTodo(input: $input) {\n    id\n    title\n    done\n    createdAt\n    doneAt\n    __typename\n  }\n}"
		}
	`, notificationTitle))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	fmt.Println("✌ ✌ ✌ Todo message sent successfully.")
}
