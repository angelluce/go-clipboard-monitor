package internal

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Notifier interface {
	Notify(result Result)
}

type ConsoleNotifier struct {
	Logger *log.Logger
}

func (c *ConsoleNotifier) Notify(result Result) {
	timestamp := time.Now().Format("15:04:05")

	fmt.Print("\033[2K\r")

	fmt.Println("╭──────────────────────────────────────────────────────────╮")
	fmt.Printf("  🔒 Contenido sensible detectado - %s\n", timestamp)
	fmt.Printf("\n  Reglas activadas: %d\n", len(result.TriggeredRules))
	for _, rule := range result.TriggeredRules {
		fmt.Printf("   • %s\n", rule)
	}
	fmt.Printf("  Contenido protegido y reemplazado\n")
	fmt.Println("╰──────────────────────────────────────────────────────────╯")

	fmt.Print("> ")

	logMessage := strings.Builder{}
	logMessage.WriteString(fmt.Sprintf("Contenido sensible detectado - Reglas: [%s]",
		strings.Join(result.TriggeredRules, ", ")))

	c.Logger.Println(logMessage.String())
}
