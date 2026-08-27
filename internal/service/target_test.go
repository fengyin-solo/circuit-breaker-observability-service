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

func TestTarget030(t *testing.T) {
	svc := evaluatorService()
	st := svc.NewRuntimeState()
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	done := make(chan error, 1)
	go func() { done <- st.SaveContext(ctx, []string{"open"}, map[string]string{"kind": "breaker"}) }()
	select {
	case err := <-done:
		if err == nil {
			t.Fatal("取消后仍保存状态")
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("取消后保存未结束")
	}
}
