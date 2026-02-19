package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

type Config struct {
	Words map[string]string `json:"words"`
}

const configPath = "replacements.json"

func main() {
	fmt.Println("\n     ✨ Hola, bienvenido a CLIPBOARD MONITOR")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("🕵️ Escaneando portapapeles...")
	fmt.Println("🛟 Usa el comando \"help\" si necesitas ayuda.")
	fmt.Println(strings.Repeat("-", 50))

	go runMonitor()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Usamos Regex para separar por espacios pero respetar lo que está entre comillas
		re := regexp.MustCompile(`"([^"]+)"|([^\s]+)`)
		matches := re.FindAllStringSubmatch(input, -1)

		var args []string
		for _, m := range matches {
			if m[1] != "" {
				args = append(args, m[1]) // Contenido entre comillas
			} else {
				args = append(args, m[2]) // Palabra simple
			}
		}

		switch args[0] {
		case "add":
			if len(args) < 3 {
				fmt.Println("❌ Uso: add \"frase buscar\" \"frase reemplazo\"")
			} else {
				agregarPalabra(args[1], args[2])
			}
		case "list":
			listarPalabras()
		case "help":
			mostrarAyuda()
		case "exit", "quit":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Printf("❓ Comando desconocido: %s\n", args[0])
		}
	}
}

func mostrarAyuda() {
	fmt.Println("\n📖 COMANDOS:")
	fmt.Println("  add \"p\" \"r\"   -> Agrega una regla")
	fmt.Println("  list          -> Muestra todas las reglas")
	fmt.Println("  exit          -> Cierra el programa")
}

func agregarPalabra(p, r string) {
	config := cargarConfig()
	config.Words[p] = r
	saveConfig(config)
	fmt.Printf("✅ Regla añadida: [%s] ➔ [%s]\n", p, r)
}

func listarPalabras() {
	config := cargarConfig()
	if len(config.Words) == 0 {
		fmt.Println("📭 No hay reglas.")
		return
	}
	fmt.Println("\n📋 REGLAS ACTUALES:")
	for k, v := range config.Words {
		fmt.Printf("  %-20s ➔  %s\n", k, v)
	}
}

func runMonitor() {
	lastContent, _ := clipboard.ReadAll()

	for {
		config := cargarConfig()
		currentContent, _ := clipboard.ReadAll()

		if currentContent != "" && currentContent != lastContent {
			modified := currentContent
			found := false
			for k, v := range config.Words {
				if strings.Contains(modified, k) {
					modified = strings.ReplaceAll(modified, k, v)
					found = true
				}
			}

			if found {
				// \r regresa al inicio de linea, \033[K borra la linea actual
				fmt.Print("\r\033[K")
				fmt.Printf("⚠️ Contenido sensible - %s\n> ", time.Now().Format("15:04:05"))
				clipboard.WriteAll(modified)
				lastContent = modified
			} else {
				lastContent = currentContent
			}
		}
		time.Sleep(1 * time.Second)
	}
}

func cargarConfig() Config {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return Config{Words: make(map[string]string)}
	}
	var c Config
	json.Unmarshal(data, &c)
	if c.Words == nil {
		c.Words = make(map[string]string)
	}
	return c
}

func saveConfig(c Config) {
	data, _ := json.MarshalIndent(c, "", "  ")
	os.WriteFile(configPath, data, 0644)
}
