package goenv

import "strconv"

func Int64(env string, required bool, dft int64) int64 {
	value, exists := parse(env, required)
	if !exists {
		return dft
	}
	result, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}
