package web

import (
	interfaces "github.com/webdelo/tradebot/interfaces/market"
	"net/http"
)

var Messages = make(chan interfaces.Kline)

// helloFromClient is a method that handles messages from the app client.
func helloFromClient(c *Client, data interface{}) {
	for {
		select {
		case message := <-Messages:
			c.send = Message{Name: "helloFromServer", Data: message}
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
