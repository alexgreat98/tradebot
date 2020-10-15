package web

import (
	"fmt"
	"github.com/webdelo/tradebot/pkg/market"
	"net/http"
)

var Messages = make(chan market.TradeDto)

// helloFromClient is a method that handles messages from the app client.
func helloFromClient(c *Client, data interface{}) {
	for {
		select {
		case message := <-Messages:
			fmt.Println("message", message.Price())
			c.send = Message{Name: "helloFromServer", Data: message.Quantity()}
			c.Write()
		}
	}
}

func Run() {
	// create router instance
	router := NewRouter()

	// handle events with messages named `helloFromClient` with handler
	// helloFromClient (from above).
	router.Handle("helloFromClient", helloFromClient)

	// handle all requests to /, upgrade to WebSocket via our router handler.
	http.Handle("/", router)

	// start server.
	http.ListenAndServe(":4000", nil)
}
