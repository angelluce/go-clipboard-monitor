package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileReader interface {
	Read(path string) (string, error)
	Write(path string, content string) error
	SupportedExtensions() []string
}

type TextFileReader struct{}

func (r *TextFileReader) Read(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error al leer archivo: %w", err)
	}
	return string(data), nil
}

func (r *TextFileReader) Write(path string, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("error al escribir archivo: %w", err)
	}
	return nil
}

func (r *TextFileReader) SupportedExtensions() []string {
	return []string{".txt", ".log"}
}

func GetReaderForFile(path string) (FileReader, error) {
	ext := strings.ToLower(filepath.Ext(path))

	textReader := &TextFileReader{}
	for _, supported := range textReader.SupportedExtensions() {
		if ext == supported {
			return textReader, nil
		}
	}

	return nil, fmt.Errorf("formato de archivo no soportado: %s", ext)
}
