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

func TestTarget029(t *testing.T) {
	svc := evaluatorService()
	st := svc.NewRuntimeState()
	st.Save([]string{"open"}, map[string]string{"kind": "breaker"})
	vals, labels := st.Load()
	vals[0] = "closed"
	labels["x"] = "y"
	again, againLabels := st.Load()
	if again[0] != "open" || againLabels["kind"] != "breaker" {
		t.Fatal("读取结果污染状态")
	}
	_ = context.Background()
}
