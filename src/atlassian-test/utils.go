package main

import "os"

func getEnv(key, empty string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return empty
	}
	return v
}
