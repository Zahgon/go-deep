package training

import (
	"text/tabwriter"
	"time"

	deep "github.com/patrikeh/go-deep"
)

// StatsPrinter prints training progress
type StatsPrinter struct {
	w *tabwriter.Writer
}

// NewStatsPrinter creates a StatsPrinter
func NewStatsPrinter() *StatsPrinter { _ = "STUB: not implemented"; return nil }

// Init initializes printer
func (p *StatsPrinter) Init(n *deep.Neural) { _ = "STUB: not implemented"; return }

// PrintProgress prints the current state of training
func (p *StatsPrinter) PrintProgress(n *deep.Neural, validation Examples, elapsed time.Duration, iteration int) {
	_ = "STUB: not implemented"
	return
}

func formatAccuracy(n *deep.Neural, validation Examples) string {
	_ = "STUB: not implemented"
	return ""
}

func accuracy(n *deep.Neural, validation Examples) float64 { _ = "STUB: not implemented"; return 0 }

func crossValidate(n *deep.Neural, validation Examples) float64 {
	_ = "STUB: not implemented"
	return 0
}
