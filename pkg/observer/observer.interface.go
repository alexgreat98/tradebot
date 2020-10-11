package observer

type Observable interface {
	//Subscribe new observer and will notify about new events
	Subscribe(subscriber Observer)

	//Unsubscribe observer and will not send new events
	Unsubscribe(subscriber Observer)
}

type Observer interface {
	//HandleEvent method for processing events from observable objects
	HandleEvent(event string, data interface{}) error

	//ObserverKey retrieve unique key (using for unsubscribe method)
	ObserverKey() string
}
