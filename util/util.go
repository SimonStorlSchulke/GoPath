package util

//clamps a float32 value from min to max
func Clamp32(val, min, max float32) float32 {
	if val > max {
		return max
	} else if val < min {
		return min
	}
	return val
}

//clamps a float64 value from min to max
func Clamp64(val, min, max float64) float64 {
	if val > max {
		return max
	} else if val < min {
		return min
	}
	return val
}
