package server

import (
	"context"
	"sync"
)

// startListeners runs every chain subscription and poller in its own
// goroutine. They stop when ctx is cancelled; the returned WaitGroup reports
// when they have all finished.
func (d *dependencies) startListeners(ctx context.Context) *sync.WaitGroup {
	var listeners sync.WaitGroup

	for _, listen := range []func(context.Context){
		d.paymentProcessor.ListenToPaymentReceivedEvent,
		d.paymentProcessor.ListenToReleaseEvent,
		d.multisig.ListenToEvents,
		d.paymentProcessorStorage.ListenToPauseEvents,
		d.paymentAutomation.PollDueTasks,
	} {
		listeners.Add(1)
		go func() {
			defer listeners.Done()
			listen(ctx)
		}()
	}

	return &listeners
}
