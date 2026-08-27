package service

import (
	"circuitbreaker/internal/config"
	"circuitbreaker/internal/store"
	"circuitbreaker/pkg/logger"
	"context"
	"testing"
	"time"
)

func evaluatorService() *Service {
	return New(store.NewMemoryStore(), logger.NewLevel(logger.LevelError), &config.Config{MaxPageSize: 100})
}

func TestTarget012(t *testing.T) {
	svc := evaluatorService()
	ctx, cancel := context.WithCancel(context.Background())
	release := make(chan struct{})
	done := make(chan error, 1)
	go func() { done <- svc.WaitSnapshotWorkers(ctx, 2, release) }()
	time.Sleep(time.Millisecond)
	cancel()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("批次未等待或未响应取消")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("批次未结束")
	}
}
