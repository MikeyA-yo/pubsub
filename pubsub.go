package pubsub

import (
	"github.com/MikeyA-yo/pubsub/pub"
	"github.com/MikeyA-yo/pubsub/sub"
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
