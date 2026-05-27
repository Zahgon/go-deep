package training

import (
	deep "github.com/patrikeh/go-deep"
)

// Trainer is a neural network trainer
type Trainer interface {
	Train(n *deep.Neural, examples, validation Examples, iterations int)
}

// OnlineTrainer is a basic, online network trainer
type OnlineTrainer struct {
	*internal
	solver    Solver
	printer   *StatsPrinter
	verbosity int
}

// NewTrainer creates a new trainer
func NewTrainer(solver Solver, verbosity int) *OnlineTrainer { _ = "STUB: not implemented"; return nil }

type internal struct {
	deltas [][]float64
}

func newTraining(layers []*deep.Layer) *internal { _ = "STUB: not implemented"; return nil }

// Train trains n
func (t *OnlineTrainer) Train(n *deep.Neural, examples, validation Examples, iterations int) {
	_ = "STUB: not implemented"
	return
}

func (t *OnlineTrainer) learn(n *deep.Neural, e Example, it int) { _ = "STUB: not implemented"; return }

func (t *OnlineTrainer) calculateDeltas(n *deep.Neural, ideal []float64) {
	_ = "STUB: not implemented"
	return
}

func (t *OnlineTrainer) update(n *deep.Neural, it int) { _ = "STUB: not implemented"; return }
