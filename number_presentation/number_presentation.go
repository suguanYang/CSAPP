package number_presentation

import "math"

func Tmax64() int64 {
	return int64(math.Pow(2, 63)) - 1
}

func Tmin64() int64 {
	return int64(math.Pow(2, 63))
}

func Umax64() uint64 {
	// can not use math.Pow here, Pow return float64 which is a signed float
	return ^uint64(0)
}

func Umin64() uint64 {
	return 0
}
