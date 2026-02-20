package internal

import (
	"time"

	"github.com/atotto/clipboard"
)

type Monitor struct {
	Engine   *Engine
	Metrics  *Metrics
	Notifier Notifier
	DryRun   bool
}

func NewMonitor(engine *Engine, metrics *Metrics, notifier Notifier, dryRun bool) *Monitor {
	return &Monitor{
		Engine:   engine,
		Metrics:  metrics,
		Notifier: notifier,
		DryRun:   dryRun,
	}
}

func (m *Monitor) Run() {
	lastContent, _ := clipboard.ReadAll()

	for {
		currentContent, _ := clipboard.ReadAll()

		if currentContent != "" && currentContent != lastContent {
			result := m.Engine.Process(currentContent)

			if len(result.TriggeredRules) > 0 {
				m.Metrics.Register(result.TriggeredRules)

				if !m.DryRun {
					clipboard.WriteAll(result.ModifiedText)
				}

				time.Sleep(100 * time.Millisecond)
				m.Notifier.Notify(result)
				lastContent = result.ModifiedText
			} else {
				lastContent = currentContent
			}
		}

		time.Sleep(1 * time.Second)
	}
}
