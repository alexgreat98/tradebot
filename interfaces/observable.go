package interfaces

type Observable interface {
	Subscribe(subscriber Observer)
	Unsubscribe(subscriber Observer)
}

type Observer interface {
	HandleEvent(event interface{}) error
	ObserverKey() string
}
