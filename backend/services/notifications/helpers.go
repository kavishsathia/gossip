package notifications

import (
	"backend/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendNotification(c *gin.Context, userId int, message string, origin int) {
	if origin == userId {
		return
	}

	// By right, this should be done individually in each endpoint
	// but this works too
	helpers.IncrementAura(userId)

	rdb, err := helpers.OpenRedis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the pubsub server"})
		return
	}

	defer rdb.Close()

	err = rdb.Publish(c, strconv.Itoa(userId), message).Err()
	if err != nil {
		panic(err)
	}
}

func SendThreadInfo(c *gin.Context, threadId int, opType string, data int) {
	rdb, err := helpers.OpenRedis()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the pubsub server"})
		return
	}

	defer rdb.Close()

	err = rdb.Publish(c, "t"+strconv.Itoa(threadId), opType+":"+strconv.Itoa(data)).Err()
	if err != nil {
		panic(err)
	}
}
