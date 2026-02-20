package main

import (
	"go-clipboard-monitor/internal"
)

func main() {
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
