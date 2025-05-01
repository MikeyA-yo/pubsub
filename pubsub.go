package pubsub

import (
	"pubsub/pub"
	"pubsub/sub"
)

type PubSub struct {
	Publisher  *pub.Pub
	Subscriber *sub.Sub
}

func NewPubSub() *PubSub {
	return &PubSub{
		Publisher:  pub.NewPub(),
		Subscriber: sub.NewSub(),
	}
}
