package config

import (
	"os"
)

func IsTest() bool {
	env_isTest := os.Getenv("IS_TEST")
	return env_isTest != ""
}

func IsTestonDocker() bool {
	env_isonDockerTest := os.Getenv("IS_TEST_ON_DOCKER")
	return env_isonDockerTest != ""
}

func DBUser() string {
	if IsTest() || IsTestonDocker() {
		return "root"
	}

	return "TODO"
}

func DBPass() string {
	if IsTest() || IsTestonDocker() {
		return "De3thM3rch"
	}
	return "TODO"
}

func DBMethod() string {
	if IsTest() {
		return "tcp(localhost:3306)"
	} else if IsTestonDocker() {
		return "tcp(db:3306)"
	}
	return "TODO"
}
