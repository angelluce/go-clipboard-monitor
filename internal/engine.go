package internal

import "strings"

type Engine struct {
	Rules map[string]string
}

type Result struct {
	OriginalText   string
	ModifiedText   string
	TriggeredRules []string
}

func NewEngine(rules map[string]string) *Engine {
	return &Engine{Rules: rules}
}

func (e *Engine) UpdateRules(rules map[string]string) {
	e.Rules = rules
}

func (e *Engine) Process(input string) Result {
	modified := input
	var triggered []string

	for k, v := range e.Rules {
		if strings.Contains(modified, k) {
			modified = strings.ReplaceAll(modified, k, v)
			triggered = append(triggered, k)
		}
	}

	return Result{
		OriginalText:   input,
		ModifiedText:   modified,
		TriggeredRules: triggered,
	}
}
