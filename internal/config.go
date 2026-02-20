package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

const configPath = "replacements.json"

type Config struct {
	Words map[string]string `json:"words"`
}

func LoadConfig() Config {
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

func AddRule(k, v string) {
	config := LoadConfig()
	config.Words[k] = v
	SaveConfig(config)
}

func ListRules() {
	config := LoadConfig()
	fmt.Println("╭──────────────────────────────────────────────────────────╮")
	fmt.Println("  📋 Reglas activas:")
	fmt.Println("")
	for k, v := range config.Words {
		fmt.Printf("  %-24s > %s\n", k, v)
	}
	fmt.Println("╰──────────────────────────────────────────────────────────╯")
}

func (m *Metrics) PrintStats() {
	fmt.Println("╭──────────────────────────────────────────────────────────╮")
	fmt.Println("  📊 Estadísticas de protección")
	fmt.Printf("\n  Total de reemplazos: %-15d", m.TotalHits)
	fmt.Println("")

	if len(m.RuleHits) == 0 {
		fmt.Println("  No hay reglas activadas aún")
		fmt.Println("╰──────────────────────────────────────────────────────────╯")
		return
	}

	fmt.Println("  Detalle por regla:")
	for k, v := range m.RuleHits {
		fmt.Printf("  • %-20s: %d veces\n", k, v)
	}
	fmt.Println("╰──────────────────────────────────────────────────────────╯")
}

func SaveConfig(c Config) {
	data, _ := json.MarshalIndent(c, "", "  ")
	os.WriteFile(configPath, data, 0644)
}
