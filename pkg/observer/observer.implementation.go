package observer

type ObservableImpl struct {
	subscribers []Observer
}

//Subscribe new observer and will notify about new events
func (o *ObservableImpl) Subscribe(subscriber Observer) {
	o.subscribers = append(o.subscribers, subscriber)
}

//Unsubscribe observer and will not send new events
func (o *ObservableImpl) Unsubscribe(subscriber Observer) {
	o.subscribers = removeFromSlice(o.subscribers, subscriber)
}

func removeFromSlice(subscribers []Observer, subscriber Observer) []Observer {
	observerListLength := len(subscribers)
	for i, observer := range subscribers {
		if subscriber.ObserverKey() == observer.ObserverKey() {
			subscribers[observerListLength-1], subscribers[i] = subscribers[i], subscribers[observerListLength-1]
			return subscribers[:observerListLength-1]
		}
	}
	return subscribers
}

//NotifySubscribers about new event and transmits the data
func (o *ObservableImpl) NotifySubscribers(event string, data interface{}) {
	for _, subscriber := range o.subscribers {
		// TODO: use named type for event detection
		go subscriber.HandleEvent(event, data)
	}
}
