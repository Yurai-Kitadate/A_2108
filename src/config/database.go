package config

import (
	"os"
)

func isTest() bool {
	env_isTest := os.Getenv("IS_TEST")
	return env_isTest == "" 
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
