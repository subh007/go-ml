package probability

// FindProbability gives the probability with given
// sample size n and m favorable event.
// It returns probability.
func FindProbability(n, m float32) float32 {
	return m / n
}
