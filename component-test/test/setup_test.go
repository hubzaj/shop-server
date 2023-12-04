package test

import (
	"context"
	"os"
	"testing"
)

var (
	ctx    context.Context
	cancel context.CancelFunc
)

func TestMain(m *testing.M) {
	ctx, cancel = context.WithCancel(context.Background())

	exitCode := m.Run()
	cancel()
	os.Exit(exitCode)
}
