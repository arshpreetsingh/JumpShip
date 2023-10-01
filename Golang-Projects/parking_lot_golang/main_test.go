package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	MainSetup()
	defer MainTearDown()
	os.Exit(m.Run())
}

func MainSetup() {
	os.Setenv("ENV", "dev")
}

func MainTearDown() {
	// flush all the env
	os.Clearenv()
}
