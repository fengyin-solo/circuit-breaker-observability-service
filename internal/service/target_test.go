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

func TestTarget010(t *testing.T) {
	svc := evaluatorService()
	labels := map[string]string{"state": "open"}
	out, err := svc.CloneEventLabels(context.Background(), labels)
	if err != nil {
		t.Fatalf("复制标签返回错误: %v", err)
	}
	out["state"] = "closed"
	if labels["state"] != "open" {
		t.Fatal("返回标签污染原状态")
	}
}
