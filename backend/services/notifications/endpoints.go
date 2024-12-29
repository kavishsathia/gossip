package notifications

import (
	"backend/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetNotifications(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		print(err.Error())
		return
	}

	defer conn.Close()

	_, userInfo, err := helpers.GetUserInfo(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rdb, err := helpers.OpenRedis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the pubsub server"})
		return
	}

	defer rdb.Close()

	pubsub := rdb.Subscribe(c, strconv.Itoa(userInfo.UserID))

	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}
}

func GetThreadInfo(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		print(err.Error())
		return
	}

	defer conn.Close()

	rdb, err := helpers.OpenRedis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the pubsub server"})
		return
	}

	defer rdb.Close()

	pubsub := rdb.Subscribe(c, "t"+c.Param("id"))

	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}
}
