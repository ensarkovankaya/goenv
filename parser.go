package goenv

import (
	"fmt"
	"os"
	"strings"
)

// Find an environment variable in
func find(env string) (string, bool) {
	for _, value := range os.Environ() {
		if strings.HasPrefix(value, env) {
			return value[len(env)+1:], true
		}
	}
	return "", false
}

func parse(env string, required bool) (string, bool) {
	value, exists := find(env)
	if !exists && required {
		panic(fmt.Sprintf("environment variable not defined: %v", env))
	}
	return value, exists
}
