package test

import (
	"context"
	"os"
	"testing"
)

var (
	cancel context.CancelFunc
)

func TestMain(m *testing.M) {
	_, cancel = context.WithCancel(context.Background())

	exitCode := m.Run()
	cancel()
	os.Exit(exitCode)
}
