package config

import (
	"os"
)

func isTest() bool {
	env_isTest := os.Getenv("IS_TEST")
	return env_isTest != ""
}

func isTestonDocker() bool {
	env_isonDockerTest := os.Getenv("IS_TEST_ON_DOCKER")
	return env_isonDockerTest != ""
}

func DBUser() string {
	if isTest() || isTestonDocker() {
		return "root"
	}

	return "TODO"
}

func DBPass() string {
	if isTest() || isTestonDocker() {
		return "De3thM3rch"
	}
	return "TODO"
}

func DBMethod() string {
	if isTest() {
		return "tcp(localhost:3306)"
	} else if isTestonDocker() {
		return "tcp(db:3306)"
	}
	return "TODO"
}
