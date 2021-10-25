package config

import (
	"os"
)

func isTest() bool {
	env_isTest := os.Getenv("IS_TEST")
	if env_isTest == "" {
		return false
	}
	return true
}

func DBUser() string {
	if (isTest()) {
		return "root"
	}

	return "TODO"
}

func DBPass() string {
	if (isTest()){
		return "De3thM3rch"
	}
	return "TODO"
}

func DBMethod() string {
	if (isTest()){
		return "tcp(localhost:3306)"
	}
	return "TODO"
}
