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

## Semantic Versioning

This project uses semantic versioning through automated GitHub Actions. Version tags are automatically created based on commit message conventions:

### Version Bump Rules
- **Major Version** (X.0.0): Include `BREAKING CHANGE` or `major:` in commit message
- **Minor Version** (0.X.0): Include `feat:` or `minor:` in commit message
- **Patch Version** (0.0.X): All other commits

### Examples
```bash
# Triggers a major version bump (X.0.0)
git commit -m "refactor: complete API redesign
BREAKING CHANGE: new API structure"

# Triggers a minor version bump (0.X.0)
git commit -m "feat: add slack integration"

# Triggers a patch version bump (0.0.X)
git commit -m "fix: handle error in teams notification"
```

## License

MIT