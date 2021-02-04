package goenv

func String(env string, required bool, dft string) string {
	value, exists := parse(env, required)
	if !exists {
		return dft
	}
	return value
}
