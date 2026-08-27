package service

import (
	"circuitbreaker/internal/config"
	"circuitbreaker/internal/store"
	"circuitbreaker/pkg/logger"
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"
)

func evaluatorService() *Service {
	return New(store.NewMemoryStore(), logger.NewLevel(logger.LevelError), &config.Config{MaxPageSize: 100})
}

func TestTarget007(t *testing.T) {
	svc := evaluatorService()
	ctx, cancel := context.WithCancel(context.Background())
	started := make(chan struct{})
	var calls atomic.Int32
	done := make(chan error, 1)
	go func() {
		done <- svc.RetrySnapshot(ctx, 3, func(callCtx context.Context) error {
			if calls.Add(1) == 1 {
				close(started)
			}
			<-callCtx.Done()
			return callCtx.Err()
		})
	}()
	<-started
	cancel()
	select {
	case err := <-done:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("取消后错误不对: %v", err)
		}
		if calls.Load() != 1 {
			t.Fatalf("取消后仍重试: %d", calls.Load())
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("取消后重试未停止")
	}
}
