package deep

// Dump is a neural network dump
type Dump struct {
	Config  *Config
	Weights [][][]float64
}

// ApplyWeights sets the weights from a three-dimensional slice
func (n *Neural) ApplyWeights(weights [][][]float64) { _ = "STUB: not implemented"; return }

// Weights returns all weights in sequence
func (n Neural) Weights() [][][]float64 { _ = "STUB: not implemented"; return nil }

// Dump generates a network dump
func (n Neural) Dump() *Dump { _ = "STUB: not implemented"; return nil }

// FromDump restores a Neural from a dump
func FromDump(dump *Dump) *Neural { _ = "STUB: not implemented"; return nil }

// Marshal marshals to JSON from network
func (n Neural) Marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Unmarshal restores network from a JSON blob
func Unmarshal(bytes []byte) (*Neural, error) { _ = "STUB: not implemented"; return nil, nil }
