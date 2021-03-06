package usvc

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/go-logr/logr"
)

// InstallSignalHandlers registers handlers for common OS signals, and closes the returned channel if they're received.
func InstallSignalHandlers(log logr.Logger) <-chan struct{} {
	stopCh := make(chan struct{})
	signalCh := make(chan os.Signal, 2)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signalCh
		log.Info("Got signal", "signal", sig)
		close(stopCh)

		sig = <-signalCh
		log.Info("Got signal", "signal", sig)
		os.Exit(1) // user is insistent
	}()

	return stopCh
}

// ChannelWrapper runs an error-returning func on a background goroutine.
// If the func returns an error, that is sent down the channel that's returned.
func ChannelWrapper(fn func() error) <-chan error {
	ch := make(chan error)
	go func() {
		defer close(ch)
		ch <- fn()
	}()
	return ch
}
