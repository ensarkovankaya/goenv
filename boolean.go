package goenv

import "strconv"

func Bool(env string, required bool, dft bool) bool {
	value, exists := parse(env, required)
	if !exists {
		return dft
	}
	result, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}
	return result
}
