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

func TestTarget025(t *testing.T) {
	svc := evaluatorService()
	want := errors.New("downstream failed")
	err := svc.PreserveHealthFailure(context.Background(), func() error { return want })
	if !errors.Is(err, want) {
		t.Fatalf("失败未回传: %v", err)
	}
}
