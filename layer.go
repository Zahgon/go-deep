package deep

// Layer is a set of neurons and corresponding activation
type Layer struct {
	Neurons []*Neuron
	A       ActivationType
}

// NewLayer creates a new layer with n nodes
func NewLayer(n int, activation ActivationType) *Layer { _ = "STUB: not implemented"; return nil }

func (l *Layer) fire() { _ = "STUB: not implemented"; return }

// Connect fully connects layer l to next, and initializes each
// synapse with the given weight function
func (l *Layer) Connect(next *Layer, weight WeightInitializer) { _ = "STUB: not implemented"; return }

// ApplyBias creates and returns a bias synapse for each neuron in l
func (l *Layer) ApplyBias(weight WeightInitializer) []*Synapse {
	_ = "STUB: not implemented"
	return nil
}

func (l Layer) String() string { _ = "STUB: not implemented"; return "" }
