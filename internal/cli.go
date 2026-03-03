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
	Scanner *Scanner
}

func NewCLI(metrics *Metrics, engine *Engine) *CLI {
	return &CLI{
		Metrics: metrics,
		Engine:  engine,
		Scanner: NewScanner(engine),
	}
}

func (cli *CLI) PrintWelcome() {
	fmt.Println(BoxTop)
	fmt.Printf("  %s✨  Hola, bienvenido a CLIPBOARD MONITOR %s %s\n", ColorPrimary, Version, ColorReset)
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
	fmt.Printf("  %sscan ruta/archivo.txt     Escanea y sanitiza archivo%s\n", ColorGreen, ColorReset)
	fmt.Printf("  %shelp                      Muestra esta ayuda%s\n", ColorGreen, ColorReset)
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

	case "scan":
		if len(args) < 2 {
			fmt.Println(BoxTop)
			fmt.Printf("  %s❌  Uso: scan ruta/archivo.txt%s\n", ColorYellow, ColorReset)
			fmt.Println(BoxBottom)
			return
		}
		cli.handleScan(args[1])

	case "help":
		cli.printHelp()

	default:
		fmt.Println(BoxTop)
		fmt.Printf("  %s❓  Comando desconocido. Usa 'help' para ver opciones.%s\n", ColorYellow, ColorReset)
		fmt.Println(BoxBottom)
	}
}

func (cli *CLI) handleScan(filePath string) {
	fmt.Println(BoxTop)
	fmt.Printf("  %s🔍 Escaneando archivo: %s%s\n", ColorBlue, filePath, ColorReset)
	fmt.Println(BoxBottom)

	result, err := cli.Scanner.ScanFile(filePath, false)
	if err != nil {
		fmt.Println(BoxTop)
		fmt.Printf("  %s❌ Error: %v%s\n", ColorYellow, err, ColorReset)
		fmt.Println(BoxBottom)
		return
	}

	fmt.Println(BoxTop)
	if len(result.TriggeredRules) > 0 {
		fmt.Printf("  %s✅  Archivo sanitizado correctamente%s\n", ColorGreen, ColorReset)
		fmt.Printf("\n  %s📊 Resumen:%s\n", ColorPrimary, ColorReset)
		fmt.Printf("  %s   • Reglas aplicadas: %d%s\n", ColorBlue, len(result.TriggeredRules), ColorReset)
		fmt.Printf("  %s   • Reemplazos realizados: %d%s\n", ColorBlue, result.ReplacedCount, ColorReset)
		fmt.Printf("  %s   • Archivo generado: %s%s\n", ColorGreen, result.OutputPath, ColorReset)
	} else {
		fmt.Printf("  %s✅ Archivo procesado%s\n", ColorGreen, ColorReset)
		fmt.Printf("  %s   • No se encontraron términos sensibles%s\n", ColorBlue, ColorReset)
	}
	fmt.Println(BoxBottom)
}
