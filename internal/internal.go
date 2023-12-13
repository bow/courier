// Copyright (c) 2022 Wibowo Arindrarto <contact@arindrarto.dev>
// SPDX-License-Identifier: BSD-3-Clause

package internal

import (
	"fmt"
	"os"
	"strings"
)

// AppName returns the application name.
func AppName() string {
	return "lens"
}

// EnvKey returns the environment variable key for configuration.
func EnvKey(key string) string {
	if key == "" {
		return envPrefix()
	}
	return fmt.Sprintf("%s_%s", envPrefix(), strings.ToUpper(strings.ReplaceAll(key, "-", "_")))
}

// Banner shows the application name as ASCII art.
func Banner() string {
	return `    __
   / /   ___   ____   _____
  / /   / _ \ / __ \ / ___/
 / /___/  __// / / /(__  )
/_____/\___//_/ /_//____/`
}

func Dedup[T comparable](values []T) []T {
	seen := make(map[T]struct{})
	nodup := make([]T, 0)

	for _, val := range values {
		if _, exists := seen[val]; exists {
			continue
		}
		seen[val] = struct{}{}
		nodup = append(nodup, val)
	}

	return nodup
}

// envPrefix returns the environment variable prefix for configuration.
func envPrefix() string {
	return strings.ToUpper(AppName())
}

func getOrExit[T any](key string, f func(string) (T, error), fallback T) T {
	var (
		err    error
		parsed = fallback
	)
	if raw := os.Getenv(EnvKey(key)); raw != "" {
		parsed, err = f(raw)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1) //nolint:revive
		}
	}
	return parsed
}
