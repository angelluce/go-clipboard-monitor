package main

import (
	"context"
	"fmt"
	"go-clipboard-monitor/internal"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := internal.AcquireLock()
	if err != nil {
		fmt.Println(internal.BoxTop)
		fmt.Printf("  %s %s %s\n", internal.ColorYellow, err, internal.ColorReset)
		fmt.Println(internal.BoxBottom)
		return
	}

	defer internal.ReleaseLock()

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Capturar Ctrl+C
	sigs := make(chan os.Signal, 1)
	signal.Notify(
		sigs,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		<-sigs
		fmt.Println("\n🛑 Cerrando Clipboard Monitor...")
		cancel()
	}()

	config := internal.LoadConfig()

	engine := internal.NewEngine(config.Words)
	metrics := internal.NewMetrics()
	logger := internal.NewLogger()
	notifier := &internal.ConsoleNotifier{Logger: logger}

	monitor := internal.NewMonitor(engine, metrics, notifier, false)

	go monitor.Run()

	cli := internal.NewCLI(metrics, engine)
	cli.Run()
}
