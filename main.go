package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func clientHandler(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[CLIENT] %s\n", string(data))
	javaConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081", http.Header{})
	if err != nil {
		log.Println(err)
		return
	}
	initialJSON, _ := json.Marshal(map[string]string{"name": c.Param("name")})
	javaConn.WriteMessage(websocket.TextMessage, initialJSON)
	javaConn.ReadMessage()
	javaConn.WriteMessage(websocket.TextMessage, data)
	_, p, err := javaConn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("[SPIGOT] %s\n", string(p))
	javaConn.Close()
	c.Writer.Write(p)
}

func main() {
	c := gin.Default()
	c.POST("/:name", clientHandler)
	c.Run(":8080")
}
