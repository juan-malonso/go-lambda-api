package env

import (
	"os"
)

func GetVar(name string, defaultValue string) string {
	value := os.Getenv(name)

	if len(value) == 0 {
		return defaultValue
	} else {
		return value
	}
}
