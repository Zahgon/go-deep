package training

import (
	deep "github.com/patrikeh/go-deep"
)

// BatchTrainer implements parallelized batch training
type BatchTrainer struct {
	*internalb
	verbosity   int
	batchSize   int
	parallelism int
	solver      Solver
	printer     *StatsPrinter
}

type internalb struct {
	deltas            [][][]float64
	partialDeltas     [][][][]float64
	accumulatedDeltas [][][]float64
	moments           [][][]float64
}

func newBatchTraining(layers []*deep.Layer, parallelism int) *internalb {
	_ = "STUB: not implemented"
	return nil
}

// NewBatchTrainer returns a BatchTrainer
func NewBatchTrainer(solver Solver, verbosity, batchSize, parallelism int) *BatchTrainer {
	_ = "STUB: not implemented"
	return nil
}

// Train trains n
func (t *BatchTrainer) Train(n *deep.Neural, examples, validation Examples, iterations int) {
	_ = "STUB: not implemented"
	return
}

func (t *BatchTrainer) calculateDeltas(n *deep.Neural, ideal []float64, wid int) {
	_ = "STUB: not implemented"
	return
}

func (t *BatchTrainer) update(n *deep.Neural, it int) { _ = "STUB: not implemented"; return }
