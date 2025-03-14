package notificator

import (
	"context"
	"log"
)

// DryRunNotifier is a notifier that does nothing but log the notification - for testing
// DryRunNotifier is a notifier that does nothing but log the notification - for testing
var _ Notifier = (*DryRunNotifier)(nil) // testing the we satisfied the interface

type DryRunNotifier struct{}

func (d *DryRunNotifier) Send(ctx context.Context, title, message string, intent Intent) error {
	log.Printf("[DRY-RUN] Would send notification: %s - %s", title, message)
	return nil
}

// Mock notifier
// Mock notifier
var _ Notifier = (*MockNotifier)(nil) // testing the we satisfied the interface

// MockNotifier is a mock implementation of the Notifier interface
type MockNotifier struct {
	Calls []struct {
		Title   string
		Message string
		Intent  Intent
	}
	ReturnError error
}

// Send stores the call arguments and returns a mock error if set
func (m *MockNotifier) Send(ctx context.Context, title, message string, intent Intent) error {
	m.Calls = append(m.Calls, struct {
		Title   string
		Message string
		Intent  Intent
	}{title, message, intent})
	return m.ReturnError
}
