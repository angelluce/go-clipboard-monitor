package internal

type Metrics struct {
	TotalHits int
	RuleHits  map[string]int
}

func NewMetrics() *Metrics {
	return &Metrics{
		RuleHits: make(map[string]int),
	}
}

func (m *Metrics) Register(rules []string) {
	m.TotalHits += len(rules)
	for _, r := range rules {
		m.RuleHits[r]++
	}
}
