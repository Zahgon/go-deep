package deep

// Mode denotes inference mode
type Mode int

const (
	// ModeDefault is unspecified mode
	ModeDefault Mode = 0
	// ModeMultiClass is for one-hot encoded classification, applies softmax output layer
	ModeMultiClass Mode = 1
	// ModeRegression is regression, applies linear output layer
	ModeRegression Mode = 2
	// ModeBinary is binary classification, applies sigmoid output layer
	ModeBinary Mode = 3
	// ModeMultiLabel is for multilabel classification, applies sigmoid output layer
	ModeMultiLabel Mode = 4
)

// OutputActivation returns activation corresponding to prediction mode
func OutputActivation(c Mode) ActivationType {
	_ = "STUB: not implemented"
	return *new(ActivationType)
}

// GetActivation returns the concrete activation given an ActivationType
func GetActivation(act ActivationType) Differentiable {
	_ = "STUB: not implemented"
	return *new(Differentiable)
}

// ActivationType is represents a neuron activation function
type ActivationType int

const (
	// ActivationNone is no activation
	ActivationNone ActivationType = 0
	// ActivationSigmoid is a sigmoid activation
	ActivationSigmoid ActivationType = 1
	// ActivationTanh is hyperbolic activation
	ActivationTanh ActivationType = 2
	// ActivationReLU is rectified linear unit activation
	ActivationReLU ActivationType = 3
	// ActivationLinear is linear activation
	ActivationLinear ActivationType = 4
	// ActivationSoftmax is a softmax activation (per layer)
	ActivationSoftmax ActivationType = 5
)

// Differentiable is an activation function and its first order derivative,
// where the latter is expressed as a function of the former for efficiency
type Differentiable interface {
	F(float64) float64
	Df(float64) float64
}

// Sigmoid is a logistic activator in the special case of a = 1
type Sigmoid struct{}

// F is Sigmoid(x)
func (a Sigmoid) F(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// Df is Sigmoid'(y), where y = Sigmoid(x)
func (a Sigmoid) Df(y float64) float64 {
	_ = "STUB: not implemented"

	// Logistic is the logistic function
	return 0
}

func Logistic(x, a float64) float64 { _ = "STUB: not implemented"; return 0 }

// Tanh is a hyperbolic activator
type Tanh struct{}

// F is Tanh(x)
func (a Tanh) F(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// Df is Tanh'(y), where y = Tanh(x)
func (a Tanh) Df(y float64) float64 { _ = "STUB: not implemented"; return 0 }

// ReLU is a rectified linear unit activator
type ReLU struct{}

// F is ReLU(x)
func (a ReLU) F(x float64) float64 { _ = "STUB: not implemented"; return 0 }

// Df is ReLU'(y), where y = ReLU(x)
func (a ReLU) Df(y float64) float64 { _ = "STUB: not implemented"; return 0 }

// Linear is a linear activator
type Linear struct{}

// F is the identity function
func (a Linear) F(x float64) float64 {
	_ = "STUB: not implemented"

	// Df is constant
	return 0
}

func (a Linear) Df(x float64) float64 { _ = "STUB: not implemented"; return 0 }
