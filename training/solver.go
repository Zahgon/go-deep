package training

// Solver implements an update rule for training a NN
type Solver interface {
	Init(size int)
	Update(value, gradient float64, iteration, idx int) float64
}

// SGD is stochastic gradient descent with nesterov/momentum
type SGD struct {
	lr       float64
	decay    float64
	momentum float64
	nesterov bool
	moments  []float64
}

// NewSGD returns a new SGD solver
func NewSGD(lr, momentum, decay float64, nesterov bool) *SGD { _ = "STUB: not implemented"; return nil }

// Init initializes vectors using number of weights in network
func (o *SGD) Init(size int) { _ = "STUB: not implemented"; return }

// Update returns the update for a given weight
func (o *SGD) Update(value, gradient float64, iteration, idx int) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Adam is an Adam solver
type Adam struct {
	lr      float64
	beta    float64
	beta2   float64
	epsilon float64

	v, m []float64
}

// NewAdam returns a new Adam solver
func NewAdam(lr, beta, beta2, epsilon float64) *Adam { _ = "STUB: not implemented"; return nil }

// Init initializes vectors using number of weights in network
func (o *Adam) Init(size int) { _ = "STUB: not implemented"; return }

// Update returns the update for a given weight
func (o *Adam) Update(value, gradient float64, t, idx int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func fparam(val, fallback float64) float64 { _ = "STUB: not implemented"; return 0 }

func iparam(val, fallback int) int { _ = "STUB: not implemented"; return 0 }
