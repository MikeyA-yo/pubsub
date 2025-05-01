# PubSub

A lightweight Go implementation of the Publish-Subscribe (PubSub) messaging pattern. This project provides a simple way to decouple publishers and subscribers, allowing for scalable and modular message passing within your applications.

## Project Structure

- `pubsub.go`: Core entry point that composes the publisher and subscriber components.
- `pub/`: Contains the publisher logic (`pub.go`).
- `sub/`: Contains the subscriber logic (`sub.go`).

## Components Overview

### Pub (Publisher)
- Responsible for sending messages to a topic.
- Uses the `Publish(topic, msg, subCtx)` method to deliver messages to all subscribers of a given topic.
- Ensures thread safety by locking the subscriber context during message delivery.

### Sub (Subscriber)
- Manages subscriptions to topics and holds channels for message delivery.
- (See `sub/sub.go` for implementation details.)

### PubSub
- Combines the Publisher and Subscriber into a single struct for easy usage.
- `NewPubSub()` creates a new instance with initialized publisher and subscriber.

## Usage Example

```go
import (
    "pubsub/pub"
    "pubsub/sub"
    "pubsub"
)

func main() {
    ps := pubsub.NewPubSub()
    // Subscribe to a topic (see sub/sub.go for details)
    // Publish a message
    ps.Publisher.Publish("topic1", "Hello, World!", ps.Subscriber)
}
```

## Getting Started

1. Clone the repository.
2. Ensure you have Go 1.23.2 or later installed.
3. Explore the `pub/` and `sub/` directories for publisher and subscriber logic.
4. Use the `PubSub` struct for easy integration.

## License

MIT