package main

import (
	"os"
	"strings"
)

type EnvConfig struct {
	GistURL string `json:"GIST_URL"`
}

func NewEnv() *EnvConfig {
	str, err := os.ReadFile(".env")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(str), "\n")
	vars := make(map[string]string)
	for _, line := range lines {
		parts := strings.Split(line, "=")
		if len(parts) == 2 {
			vars[parts[0]] = parts[1]
		}
	}

	return &EnvConfig{
		GistURL: vars["GIST_URL"],
	}
}

var Env *EnvConfig = NewEnv()
