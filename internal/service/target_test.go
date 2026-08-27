package service

import (
	"circuitbreaker/internal/config"
	"circuitbreaker/internal/store"
	"circuitbreaker/pkg/logger"
	"context"
	"errors"
	"testing"
)

func evaluatorService() *Service {
	return New(store.NewMemoryStore(), logger.NewLevel(logger.LevelError), &config.Config{MaxPageSize: 100})
}

func TestTarget019(t *testing.T) {
	svc := evaluatorService()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	out, err := svc.TransitionRecovery(ctx, "closed", "open")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("取消后状态仍被提交: %q %v", out, err)
	}
}
