package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type Account struct {
	Id        int      `json:"id"`
	Email     string   `json:"email"`
	Fname     *string  `json:"fname"`
	Sname     *string  `json:"sname"`
	Phone     *string  `json:"phone"`
	Sex       string   `json:"sex"`
	Birth     int      `json:"birth"`
	Country   *string  `json:"country"`
	City      *string  `json:"city"`
	Joined    int      `json:"joined"`
	Status    *string  `json:"status"`
	Interests []string `json:"interests"`
	Premium   struct {
		Start  int `json:"start"`
		Finish int `json:"finish"`
	} `json:"premium"`
	Likes []struct {
		Id int `json:"id"`
		Ts int `json:"ts"`
	} `json:"likes"`
}

var (
	client *redis.Client
	db     map[string][]Account
)

func SetUpRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("Redis client:", client)
}

func LoadData() {
	dirname := "/tmp/data/data"
	fmt.Println("dirname", dirname)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range files {
		filename := filepath.Join(dirname, info.Name())
		fmt.Println("filename", filename)

		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		if err := json.Unmarshal(bytes, &db); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("db[\"accounts\"][0]: %#v\n", db["accounts"][0])
	fmt.Println("len(db[\"accounts\"])", len(db["accounts"]))
}

func DumpHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, I'm Web Server!"))
}

func main() {
	SetUpRedis()
	LoadData()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":80")
}
