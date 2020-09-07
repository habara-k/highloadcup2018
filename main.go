package main

import (
	"fmt"
	"net/http"

	"github.com/go-redis/redis"
)

func DumpHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, I'm Web Server!"))
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Redis client:", client)

	http.HandleFunc("/", DumpHandler)
	http.ListenAndServe(":80", nil)
}
