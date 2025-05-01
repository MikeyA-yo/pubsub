package pub

import (
	"pubsub/sub"
)

type Pub struct {
}

func NewPub() *Pub {
	return &Pub{}
}

func (p *Pub) Publish(topic, msg string, subCtx *sub.Sub) {
	subCtx.Mu.Lock()
	defer subCtx.Mu.Unlock()
	currentTopic := subCtx.Topics[topic]
	for _, ch := range currentTopic {
		ch <- msg
	}
}
