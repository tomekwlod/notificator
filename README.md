# Notificator

A flexible notification service that supports multiple notification channels with a unified interface.

## Installation

```bash
go get github.com/tomekwlod/notificator
```

## Features

- Multi-channel notification support
- Microsoft Teams integration
- Extensible interface for adding new notification channels
- Support for different message intents (Primary, Info, Warn, Error)

## Usage

```go
package main

import (
    "context"
    "github.com/tomekwlod/notificator"
    "github.com/tomekwlod/notificator/teams"
)

func main() {
    // Create a new multi-notifier
    multi := notificator.NewMultiNotifier()

    // Create and register a Teams notifier
    teamsNotifier := teams.New("MyApp", "your-webhook-url")
    multi.RegisterChannel("teams", teamsNotifier)

    // Send a notification
    ctx := context.Background()
    multi.Send(ctx, "Test Title", "Test Message", notificator.IntentInfo)
}
```

## License

MIT