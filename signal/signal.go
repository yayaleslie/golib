package signal

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var signalDef = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGABRT,
	syscall.SIGKILL,
	syscall.SIGTERM,
}

func WaitNotify(signs ...os.Signal) {
	var signalNotify []os.Signal
	if len(signs) > 0 {
		signalNotify = signs
	} else {
		signalNotify = signalDef
	}

	exit := make(chan os.Signal)
	signal.Notify(exit, signalNotify...)

	// catch exit signal
	sign := <-exit
	log.Printf("stop by exit signal '%s'", sign)
}
