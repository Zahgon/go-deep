package deep

// Neural is a neural network
type Neural struct {
	Layers []*Layer
	Biases [][]*Synapse
	Config *Config
}

// Config defines the network topology, activations, losses etc
type Config struct {
	// Number of inputs
	Inputs int
	// Defines topology:
	// For instance, [5 3 3] signifies a network with two hidden layers
	// containing 5 and 3 nodes respectively, followed an output layer
	// containing 3 nodes.
	Layout []int
	// Activation functions: {ActivationTanh, ActivationReLU, ActivationSigmoid}
	Activation ActivationType
	// Solver modes: {ModeRegression, ModeBinary, ModeMultiClass, ModeMultiLabel}
	Mode Mode
	// Initializer for weights: {NewNormal(σ, μ), NewUniform(σ, μ)}
	Weight WeightInitializer `json:"-"`
	// Loss functions: {LossCrossEntropy, LossBinaryCrossEntropy, LossMeanSquared}
	Loss LossType
	// Apply bias nodes
	Bias bool
}

// NewNeural returns a new neural network
func NewNeural(c *Config) *Neural { _ = "STUB: not implemented"; return nil }

func initializeLayers(c *Config) []*Layer { _ = "STUB: not implemented"; return nil }

func (n *Neural) fire() { _ = "STUB: not implemented"; return }

// Forward computes a forward pass
func (n *Neural) Forward(input []float64) error { _ = "STUB: not implemented"; return nil }

// Predict computes a forward pass and returns a prediction
func (n *Neural) Predict(input []float64) []float64 { _ = "STUB: not implemented"; return nil }

// NumWeights returns the number of weights in the network
func (n *Neural) NumWeights() (num int) { _ = "STUB: not implemented"; return 0 }

func (n *Neural) String() string { _ = "STUB: not implemented"; return "" }
