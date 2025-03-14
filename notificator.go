package notificator

import (
	"context"
	"errors"
	"fmt"
)

type (
	Intent string
)

const (
	IntentPrimary Intent = ""
	IntentInfo    Intent = "info"
	IntentWarn    Intent = "warn"
	IntentError   Intent = "error"
)

type Notifier interface {
	Send(ctx context.Context, title, message string, intent Intent) error
}

type MultiNotifier interface {
	Notifier // it will simply broadcast to all channels
	Get(channel string) (Notifier, error)
	List() map[string]Notifier
	Broadcast(ctx context.Context, title, message string, intent Intent) error
}

var ErrChannelNotFound = errors.New("channel not found")

type multiNotifier struct {
	channels map[string]Notifier
}

func NewMultiNotifier() *multiNotifier {
	return &multiNotifier{
		channels: make(map[string]Notifier),
	}
}

func (m *multiNotifier) RegisterChannel(channel string, notifier Notifier) {
	m.channels[channel] = notifier
}

func (m *multiNotifier) Get(channel string) (Notifier, error) {
	notifier, ok := m.channels[channel]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrChannelNotFound, channel)
	}
	return notifier, nil
}

func (m *multiNotifier) Broadcast(ctx context.Context, title, message string, intent Intent) error {
	for _, notifier := range m.channels {
		if err := notifier.Send(ctx, title, message, intent); err != nil {
			return err // Stop on first failure
		}
	}
	return nil
}

func (m *multiNotifier) List() map[string]Notifier {
	return m.channels
}

func (m *multiNotifier) Send(ctx context.Context, title, message string, intent Intent) error {
	return m.Broadcast(ctx, title, message, intent)
}
