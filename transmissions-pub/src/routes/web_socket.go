package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"log"
	"os"
	"transmissions-pub/lib"
)

var wb = lib.WbUpgrader

func WebSocket(context *gin.Context) {
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))

	if err != nil {
		return
	}

	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	ch, err := conn.Channel()

	if err != nil {
		return
	}

	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			return
		}
	}(ch)

	ws, err := wb.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			return
		}
	}(ws)

	msgs, err := ch.Consume(
		"transmissions-status",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	for msg := range msgs {
		err := ws.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			return
		}
	}
}
