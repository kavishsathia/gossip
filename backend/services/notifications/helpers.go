package notifications

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SendNotification(c *gin.Context, userId int, message string, origin int) {
  if origin == userId {
	return 
}
  rdb := redis.NewClient(&redis.Options{
	  Addr:	  "localhost:6379",
	  Password: "", 
	  DB:		  0,  
  })

  defer rdb.Close()

  err := rdb.Publish(c, strconv.Itoa(userId), message).Err()
  if err != nil {
    panic(err)
  }
}


func SendThreadInfo(c *gin.Context, threadId int, opType string, data int) {
  rdb := redis.NewClient(&redis.Options{
	  Addr:	  "localhost:6379",
	  Password: "", 
	  DB:		  0,  
  })

  defer rdb.Close()

  err := rdb.Publish(c, "t" + strconv.Itoa(threadId), opType + ":" + strconv.Itoa(data)).Err()
  if err != nil {
    panic(err)
  }
}
