package training

// Example is an input-target pair
type Example struct {
	Input    []float64
	Response []float64
}

// Examples is a set of input-output pairs
type Examples []Example

// Shuffle shuffles slice in-place
func (e Examples) Shuffle() { _ = "STUB: not implemented"; return }

// Split assigns each element to two new slices
// according to probability p
func (e Examples) Split(p float64) (first, second Examples) {
	_ = "STUB: not implemented"
	return *new(Examples), *new(Examples)
}

// SplitSize splits slice into parts of size size
func (e Examples) SplitSize(size int) []Examples { _ = "STUB: not implemented"; return nil }

// SplitN splits slice into n parts
func (e Examples) SplitN(n int) []Examples { _ = "STUB: not implemented"; return nil }

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }
