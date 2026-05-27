package deep

// A WeightInitializer returns a (random) weight
type WeightInitializer func() float64

// NewUniform returns a uniform weight generator
func NewUniform(stdDev, mean float64) WeightInitializer {
	_ = "STUB: not implemented"
	return *new(WeightInitializer)
}

// Uniform samples a value from u(mean-stdDev/2,mean+stdDev/2)
func Uniform(stdDev, mean float64) float64 { _ = "STUB: not implemented"; return 0 }

// NewNormal returns a normal weight generator
func NewNormal(stdDev, mean float64) WeightInitializer {
	_ = "STUB: not implemented"
	return *new(WeightInitializer)
}

// Normal samples a value from N(μ, σ)
func Normal(stdDev, mean float64) float64 { _ = "STUB: not implemented"; return 0 }
