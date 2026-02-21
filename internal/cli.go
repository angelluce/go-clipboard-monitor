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
	fmt.Println(BoxTop)
	fmt.Printf("  %s✨  Hola, bienvenido a CLIPBOARD MONITOR%s\n", ColorPrimary, ColorReset)
	fmt.Printf("\n  %s📋 Monitoreando el portapapeles...%s\n", ColorYellow, ColorReset)
	fmt.Printf("  %s💡 Escribe 'help' para ver todos los comandos%s\n", ColorBlue, ColorReset)
	fmt.Println(BoxBottom)
}

func (cli *CLI) printHelp() {
	fmt.Println(BoxTop)
	fmt.Printf("  %s📚 Comandos disponibles:%s\n", ColorPrimary, ColorReset)
	fmt.Printf("\n  %sadd \"texto\" \"reemplazo\"   Añade regla de protección%s\n", ColorGreen, ColorReset)
	fmt.Printf("  %slist                      Muestra reglas actuales%s\n", ColorGreen, ColorReset)
	fmt.Printf("  %sstats                     Muestra estadísticas%s\n", ColorGreen, ColorReset)
	fmt.Printf("  %shelp                      Muestra esta ayuda%s\n", ColorGreen, ColorReset)
	fmt.Printf("  %sexit                      Cierra el programa%s\n", ColorGreen, ColorReset)
	fmt.Println(BoxBottom)
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

		cli.processCommand(strings.TrimSpace(scanner.Text()))
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
			fmt.Println(BoxTop)
			fmt.Printf("  %s❌ Uso: add \"buscar\" \"reemplazo\"%s\n", ColorYellow, ColorReset)
			fmt.Println(BoxBottom)
			return
		}
		AddRule(args[1], args[2])
		config := LoadConfig()
		cli.Engine.UpdateRules(config.Words)
		fmt.Println(BoxTop)
		fmt.Printf("  %s✅  Regla añadida y activada correctamente.%s\n", ColorGreen, ColorReset)
		fmt.Println(BoxBottom)

	case "list":
		ListRules()

	case "stats":
		cli.Metrics.PrintStats()

	case "help":
		cli.printHelp()

	case "exit", "quit":
		fmt.Println(BoxTop)
		fmt.Printf("  %s👋 ¡Hasta luego! Saliendo...%s\n", ColorPrimary, ColorReset)
		fmt.Println(BoxBottom)
		os.Exit(0)

	default:
		fmt.Println(BoxTop)
		fmt.Printf("  %s❓  Comando desconocido. Usa 'help' para ver opciones.%s\n", ColorYellow, ColorReset)
		fmt.Println(BoxBottom)
	}
}
