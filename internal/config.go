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
	fmt.Println(BoxTop)
	fmt.Printf("  %s📋 Reglas activas:%s\n", ColorPrimary, ColorReset)
	fmt.Println("")
	for k, v := range config.Words {
		fmt.Printf("  %s%-24s > %s%s\n", ColorGreen, k, v, ColorReset)
	}
	fmt.Println(BoxBottom)
}

func (m *Metrics) PrintStats() {
	fmt.Println(BoxTop)
	fmt.Printf("  %s📊 Estadísticas de protección%s\n", ColorPrimary, ColorReset)
	fmt.Printf("\n  %sTotal de reemplazos: %-15d%s", ColorGreen, m.TotalHits, ColorReset)
	fmt.Println("")

	if len(m.RuleHits) == 0 {
		fmt.Printf("  %sNo hay reglas activadas aún%s\n", ColorYellow, ColorReset)
		fmt.Println(BoxBottom)
		return
	}

	fmt.Println("  Detalle por regla:")
	for k, v := range m.RuleHits {
		fmt.Printf("  %s• %-20s: %d veces%s\n", ColorGreen, k, v, ColorReset)
	}
	fmt.Println(BoxBottom)
}

func SaveConfig(c Config) {
	data, _ := json.MarshalIndent(c, "", "  ")
	os.WriteFile(configPath, data, 0644)
}
