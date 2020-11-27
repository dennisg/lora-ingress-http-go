package domain




func NewMessageBus(topic string) (MessageBus, error) {
	return newPubSubMessageBus(topic)
}
