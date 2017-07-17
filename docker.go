package main

import(
	"os"
)

func getDockerID() string{
	id, exists := os.LookupEnv("Docker_ID")
}

