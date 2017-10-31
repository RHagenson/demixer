package utils

// FloatsApproxEqual returns whether the two float64 values (a and b)
// are within threshold distance from one another
func FloatsApproxEqual(a, b, threshold float64) bool {
	if (a-b) < threshold && (b-a) < threshold {
		return true
	}
	return false
}
