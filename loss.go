package deep

// GetLoss returns a loss function given a LossType
func GetLoss(loss LossType) Loss { _ = "STUB: not implemented"; return *new(Loss) }

// LossType represents a loss function
type LossType int

func (l LossType) String() string { _ = "STUB: not implemented"; return "" }

const (
	// LossNone signifies unspecified loss
	LossNone LossType = 0
	// LossCrossEntropy is cross entropy loss
	LossCrossEntropy LossType = 1
	// LossBinaryCrossEntropy is the special case of binary cross entropy loss
	LossBinaryCrossEntropy LossType = 2
	// LossMeanSquared is MSE
	LossMeanSquared LossType = 3
)

// Loss is satisfied by loss functions
type Loss interface {
	F(estimate, ideal [][]float64) float64
	Df(estimate, ideal, activation float64) float64
}

// CrossEntropy is CE loss
type CrossEntropy struct{}

// F is CE(...)
func (l CrossEntropy) F(estimate, ideal [][]float64) float64 { _ = "STUB: not implemented"; return 0 }

// Df is CE'(...)
func (l CrossEntropy) Df(estimate, ideal, activation float64) float64 {
	_ = "STUB: not implemented"
	return 0

	// BinaryCrossEntropy is binary CE loss
}

type BinaryCrossEntropy struct{}

// F is CE(...)
func (l BinaryCrossEntropy) F(estimate, ideal [][]float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

// Df is CE'(...)
func (l BinaryCrossEntropy) Df(estimate, ideal, activation float64) float64 {
	_ = "STUB: not implemented"
	return 0

	// MeanSquared in MSE loss
}

type MeanSquared struct{}

// F is MSE(...)
func (l MeanSquared) F(estimate, ideal [][]float64) float64 { _ = "STUB: not implemented"; return 0 }

// Df is MSE'(...)
func (l MeanSquared) Df(estimate, ideal, activation float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
