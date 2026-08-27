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

func TestTarget014(t *testing.T) {
	svc := evaluatorService()
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan string)
	done := make(chan error, 1)
	go func() { done <- svc.SendBreakerEvent(ctx, out, "event") }()
	cancel()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("取消后事件发送仍报告成功")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("取消后事件发送仍阻塞")
	}
}
