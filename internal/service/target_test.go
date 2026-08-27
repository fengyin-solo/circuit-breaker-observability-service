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

func TestTarget026(t *testing.T) {
	svc := evaluatorService()
	st := svc.NewRuntimeState()
	vals := []string{"open", "probe"}
	labels := map[string]string{"kind": "breaker"}
	st.Save(vals, labels)
	got, gotLabels := st.Load()
	got[0] = "closed"
	gotLabels["kind"] = "changed"
	again, againLabels := st.Load()
	if again[0] != "open" || againLabels["kind"] != "breaker" {
		t.Fatal("状态快照被外部输入污染")
	}
	_ = context.Background()
}
