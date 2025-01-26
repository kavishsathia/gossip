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

	// Get the Redis connection
	rdb, err := helpers.OpenRedis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the pubsub server"})
		return
	}

	defer rdb.Close()

	// Get the PubSub connection
	pubsub := rdb.Subscribe(c, strconv.Itoa(userInfo.UserID))

	defer pubsub.Close()

	ch := pubsub.Channel()

	// When there is a message from the pubsub connection, we send it to the
	// user
	for msg := range ch {
		conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}
}

// The following function was made as an additional feature, but
// after implementation, I felt like it wasn't really necessary
// and only introduced new bugs, so I have kept it here for the time
// being but it is not actively being used.
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
