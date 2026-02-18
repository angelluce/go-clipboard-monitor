package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/atotto/clipboard"
)

type Config struct {
	Words map[string]string `json:"words"`
}

const configPath = "replacements.json"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			addCmd := flag.NewFlagSet("add", flag.ExitOnError)
			p := addCmd.String("p", "", "Palabra sensible")
			r := addCmd.String("r", "", "Reemplazo")
			addCmd.Parse(os.Args[2:])

			if *p == "" || *r == "" {
				fmt.Println("\n❌ Error: Faltan parámetros.")
				fmt.Println("Uso: .\\clipboard_monitor.exe add -p \"palabra\" -r \"reemplazo\"")
				return
			}
			agregarPalabra(*p, *r)
			return

		case "list":
			listarPalabras()
			return

		case "help", "-h", "--help":
			mostrarAyuda()
			return

		default:
			fmt.Printf("\n❓ Comando desconocido: %s\n", os.Args[1])
			mostrarAyuda()
			return
		}
	}

	runMonitor()
}

func mostrarAyuda() {
	fmt.Println("\n")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("     🛡️  CLIPBOARD MONITOR - GUÍA")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("\n1. Iniciar el monitor (vigila el portapapeles):")
	fmt.Println("   .\\clipboard_monitor.exe")
	fmt.Println("\n2. Agregar una nueva palabra:")
	fmt.Println("   .\\clipboard_monitor.exe add -p \"mi_secreto\" -r \"[CENSURADO]\"")
	fmt.Println("\n3. Ver todas las reglas guardadas:")
	fmt.Println("   .\\clipboard_monitor.exe list")
	fmt.Println("\n4. Ver esta ayuda:")
	fmt.Println("   .\\clipboard_monitor.exe help")
	fmt.Println(strings.Repeat("-", 50))
}

func agregarPalabra(p, r string) {
	config := cargarConfig()
	config.Words[p] = r
	saveConfig(config)
	fmt.Printf("\n✅ REGLA REGISTRADA: [%s] -> [%s]\n", p, r)
}

func listarPalabras() {
	config := cargarConfig()
	if len(config.Words) == 0 {
		fmt.Println("\n📭 No hay reglas registradas. Usa el comando 'add'.")
		return
	}
	fmt.Println("\n📋 REGLAS ACTUALES:")
	fmt.Println(strings.Repeat("-", 50))
	for k, v := range config.Words {
		fmt.Printf("%-25s -> %-20s\n", k, v)
	}
}

func runMonitor() {
	fmt.Println("\n     ✨  Hola, bienvenido a CLIPBOARD MONITOR")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("🕵️ Escaneando portapapeles... (Ctrl+C para salir)")
	fmt.Println("Puedes agregar palabras en otra terminal con 'add'")

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
				fmt.Printf("⚠️ [%s] Contenido sensible censurado.\n", time.Now().Format("15:04:05"))
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
