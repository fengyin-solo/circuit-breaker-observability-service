package service

import (
	"circuitbreaker/internal/config"
	"circuitbreaker/internal/store"
	"circuitbreaker/pkg/logger"
	"context"
	"errors"
	"testing"
	"time"
)

func evaluatorService() *Service {
	return New(store.NewMemoryStore(), logger.NewLevel(logger.LevelError), &config.Config{MaxPageSize: 100})
}

func TestTarget002(t *testing.T) {
	svc := evaluatorService()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	entries := []string{"open", "probe"}
	out, err := svc.RefreshHealthWindow(ctx, entries)
	if err == nil || !errors.Is(err, context.Canceled) {
		t.Fatalf("取消后未返回取消状态: %v", err)
	}
	if out != nil {
		t.Fatal("取消后仍返回结果")
	}
	_ = time.Now()
}
