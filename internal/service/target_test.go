package service

import (
	"circuitbreaker/internal/config"
	"circuitbreaker/internal/store"
	"circuitbreaker/pkg/logger"
	"context"
	"testing"
)

func evaluatorService() *Service {
	return New(store.NewMemoryStore(), logger.NewLevel(logger.LevelError), &config.Config{MaxPageSize: 100})
}

func TestTarget022(t *testing.T) {
	svc := evaluatorService()
	done := false
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_ = svc.CleanupHealth(ctx, &done)
	if !done {
		t.Fatal("取消后没有执行清理")
	}
}
