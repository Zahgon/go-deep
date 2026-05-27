package deep

// Neuron is a neural network node
type Neuron struct {
	A     ActivationType `json:"-"`
	In    []*Synapse
	Out   []*Synapse
	Value float64 `json:"-"`
}

// NewNeuron returns a neuron with the given activation
func NewNeuron(activation ActivationType) *Neuron { _ = "STUB: not implemented"; return nil }

func (n *Neuron) fire() { _ = "STUB: not implemented"; return }

// Activate applies the neurons activation
func (n *Neuron) Activate(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// DActivate applies the derivative of the neurons activation
func (n *Neuron) DActivate(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// Synapse is an edge between neurons
type Synapse struct {
	Weight  float64
	In, Out float64 `json:"-"`
	IsBias  bool
}

// NewSynapse returns a synapse with the specified initialized weight
func NewSynapse(weight float64) *Synapse { _ = "STUB: not implemented"; return nil }

func (s *Synapse) fire(value float64) { _ = "STUB: not implemented"; return }
