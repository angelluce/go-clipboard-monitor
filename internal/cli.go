package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type CLI struct {
	Metrics *Metrics
	Engine  *Engine
}

func NewCLI(metrics *Metrics, engine *Engine) *CLI {
	return &CLI{
		Metrics: metrics,
		Engine:  engine,
	}
}

func (cli *CLI) PrintWelcome() {
	fmt.Println("╭──────────────────────────────────────────────────────────╮")
	fmt.Println("  ✨ Hola, bienvenido a CLIPBOARD MONITOR")
	fmt.Println("\n  📋 Monitoreando el portapapeles...")
	fmt.Println("  💡 Escribe 'help' para ver todos los comandos")
	fmt.Println("╰──────────────────────────────────────────────────────────╯")
}

func (cli *CLI) printHelp() {
	fmt.Println("╭──────────────────────────────────────────────────────────╮")
	fmt.Println("  📚 Comandos disponibles:")
	fmt.Println("\n  add \"texto\" \"reemplazo\"   Añade regla de protección")
	fmt.Println("  list                      Muestra reglas actuales")
	fmt.Println("  stats                     Muestra estadísticas")
	fmt.Println("  help                      Muestra esta ayuda")
	fmt.Println("  exit                      Cierra el programa")
	fmt.Println("╰──────────────────────────────────────────────────────────╯")
}

func (cli *CLI) Run() {
	cli.PrintWelcome()

	scanner := bufio.NewScanner(os.Stdin)
	prompt := "> "

	for {
		fmt.Print(prompt)

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		cli.processCommand(input)
	}
}

func (cli *CLI) processCommand(input string) {
	re := regexp.MustCompile(`"([^"]+)"|([^\s]+)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var args []string
	for _, m := range matches {
		if m[1] != "" {
			args = append(args, m[1])
		} else {
			args = append(args, m[2])
		}
	}

	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 3 {
			fmt.Println("╭──────────────────────────────────────────────────────────╮")
			fmt.Println("  ❌ Uso: add \"buscar\" \"reemplazo\"")
			fmt.Println("╰──────────────────────────────────────────────────────────╯")
			return
		}
		AddRule(args[1], args[2])
		config := LoadConfig()
		cli.Engine.UpdateRules(config.Words)
		fmt.Println("╭──────────────────────────────────────────────────────────╮")
		fmt.Println("  ✅ Regla añadida y activada correctamente.")
		fmt.Println("╰──────────────────────────────────────────────────────────╯")

	case "list":
		ListRules()

	case "stats":
		cli.Metrics.PrintStats()

	case "help":
		cli.printHelp()

	case "exit", "quit":
		fmt.Println("╭──────────────────────────────────────────────────────────╮")
		fmt.Println("  👋 ¡Hasta luego! Saliendo...")
		fmt.Println("╰──────────────────────────────────────────────────────────╯")
		os.Exit(0)

	default:
		fmt.Println("╭──────────────────────────────────────────────────────────╮")
		fmt.Println("  ❓ Comando desconocido. Usa 'help' para ver opciones.")
		fmt.Println("╰──────────────────────────────────────────────────────────╯")
	}
}
