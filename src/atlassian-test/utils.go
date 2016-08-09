package main

import "os"

// getEnv getting environment variable or returning default value
func getEnv(key, empty string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return empty
	}
	return v
}
