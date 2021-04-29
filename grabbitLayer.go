package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "github.com/streadway/amqp"
)


func main() {
    queueDispecther := NewQueueDispecther()
    // queueDispecther := queueDispecther()
    router := gin.Default()
    router.POST("/", func(c *gin.Context) {
        requestSchema := RequestSchema{}
        c.BindJSON(&requestSchema)
        queue := c.DefaultQuery("queue", "grabbit_default")
        queueRequest := QueueRequest{requestSchema, queue}
        queueDispecther.Input <- queueRequest
        // queueDispecther <- queueRequest
        // go sendRequestFromRequestSchema(requestSchema)
        c.String(http.StatusOK, requestSchema.Url)
    })

    router.Run(":88")
}