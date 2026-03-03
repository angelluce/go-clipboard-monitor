package internal

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Scanner struct {
	Engine *Engine
}

type ScanResult struct {
	OriginalPath   string
	OutputPath     string
	TriggeredRules []string
	ReplacedCount  int
}

func NewScanner(engine *Engine) *Scanner {
	return &Scanner{Engine: engine}
}

func (s *Scanner) ScanFile(inputPath string, overwrite bool) (*ScanResult, error) {
	// Obtener el lector apropiado
	reader, err := GetReaderForFile(inputPath)
	if err != nil {
		return nil, err
	}

	// Leer el contenido del archivo
	content, err := reader.Read(inputPath)
	if err != nil {
		return nil, err
	}

	// Procesar el contenido con el engine
	result := s.Engine.Process(content)

	// Determinar la ruta de salida
	outputPath := inputPath
	if !overwrite {
		outputPath = s.generateSanitizedPath(inputPath)
	}

	// Escribir el archivo procesado
	err = reader.Write(outputPath, result.ModifiedText)
	if err != nil {
		return nil, err
	}

	// Contar reemplazos
	replacedCount := 0
	for _, rule := range result.TriggeredRules {
		replacedCount += strings.Count(content, rule)
	}

	return &ScanResult{
		OriginalPath:   inputPath,
		OutputPath:     outputPath,
		TriggeredRules: result.TriggeredRules,
		ReplacedCount:  replacedCount,
	}, nil
}

func (s *Scanner) generateSanitizedPath(originalPath string) string {
	ext := filepath.Ext(originalPath)
	nameWithoutExt := strings.TrimSuffix(originalPath, ext)
	return fmt.Sprintf("%s_sanitized%s", nameWithoutExt, ext)
}
