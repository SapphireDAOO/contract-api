package server

import (
	"context"
	"log/slog"
	"runtime/debug"
	"sync"
)

// startListeners runs every chain subscription and poller in its own
// goroutine. They stop when ctx is cancelled; the returned WaitGroup reports
// when they have all finished.
func (d *dependencies) startListeners(ctx context.Context) *sync.WaitGroup {
	var listeners sync.WaitGroup

	for name, listen := range map[string]func(context.Context){
		"payment received":   d.paymentProcessor.ListenToPaymentReceivedEvent,
		"payment released":   d.paymentProcessor.ListenToReleaseEvent,
		"multisig":           d.multisig.ListenToEvents,
		"storage pause":      d.paymentProcessorStorage.ListenToPauseEvents,
		"pause expiry":       d.paymentProcessorStorage.WatchEmergencyPause,
		"payment automation": d.paymentAutomation.PollDueTasks,
	} {
		listeners.Go(func() {
			defer func() {
				if r := recover(); r != nil {
					slog.Error("listener panicked and stopped",
						"listener", name, "panic", r, "stack", string(debug.Stack()))
				}
			}()

			listen(ctx)
			slog.Info("listener stopped", "listener", name)
		})
	}

	return &listeners
}
