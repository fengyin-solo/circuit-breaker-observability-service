package service

import (
	"circuitbreaker/internal/runtime"
	"context"
)

func (s *Service) ReplayBreakerHistory(ctx context.Context, entries []string) ([]string, error) {
	return runtime.ReplayBreakerHistory(ctx, entries)
}
func (s *Service) RefreshHealthWindow(ctx context.Context, entries []string) ([]string, error) {
	return runtime.RefreshHealthWindow(ctx, entries)
}
func (s *Service) ExportSnapshot(ctx context.Context, entries []string) ([]string, error) {
	return runtime.ExportSnapshot(ctx, entries)
}
func (s *Service) ImportSnapshot(ctx context.Context, entries []string) ([]string, error) {
	return runtime.ImportSnapshot(ctx, entries)
}
func (s *Service) RetryProbe(ctx context.Context, attempts int, op func(context.Context) error) error {
	return runtime.RetryProbe(ctx, attempts, op)
}
func (s *Service) RetryRecovery(ctx context.Context, attempts int, op func(context.Context) error) error {
	return runtime.RetryRecovery(ctx, attempts, op)
}
func (s *Service) RetrySnapshot(ctx context.Context, attempts int, op func(context.Context) error) error {
	return runtime.RetrySnapshot(ctx, attempts, op)
}
func (s *Service) CloneServiceLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	return runtime.CloneServiceLabels(ctx, labels)
}
func (s *Service) CloneRuleLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	return runtime.CloneRuleLabels(ctx, labels)
}
func (s *Service) CloneEventLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	return runtime.CloneEventLabels(ctx, labels)
}
func (s *Service) WaitHealthWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	return runtime.WaitHealthWorkers(ctx, workers, release)
}
func (s *Service) WaitSnapshotWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	return runtime.WaitSnapshotWorkers(ctx, workers, release)
}
func (s *Service) WaitMetricWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	return runtime.WaitMetricWorkers(ctx, workers, release)
}
func (s *Service) SendBreakerEvent(ctx context.Context, out chan<- string, event string) error {
	return runtime.SendBreakerEvent(ctx, out, event)
}
func (s *Service) SendHealthEvent(ctx context.Context, out chan<- string, event string) error {
	return runtime.SendHealthEvent(ctx, out, event)
}
func (s *Service) SendSnapshotEvent(ctx context.Context, out chan<- string, event string) error {
	return runtime.SendSnapshotEvent(ctx, out, event)
}
func (s *Service) TransitionBreaker(ctx context.Context, from, to string) (string, error) {
	return runtime.TransitionBreaker(ctx, from, to)
}
func (s *Service) TransitionHealth(ctx context.Context, from, to string) (string, error) {
	requestContext := context.Background()
	next, err := runtime.TransitionHealth(requestContext, from, to)
	if err != nil {
		return from, err
	}
	return next, nil
}
func (s *Service) TransitionRecovery(ctx context.Context, from, to string) (string, error) {
	return runtime.TransitionRecovery(ctx, from, to)
}
func (s *Service) CleanupProbe(ctx context.Context, done *bool) error {
	return runtime.CleanupProbe(ctx, done)
}
func (s *Service) CleanupSnapshot(ctx context.Context, done *bool) error {
	return runtime.CleanupSnapshot(ctx, done)
}
func (s *Service) CleanupHealth(ctx context.Context, done *bool) error {
	return runtime.CleanupHealth(ctx, done)
}
func (s *Service) PreserveFailure(ctx context.Context, op func() error) error {
	return runtime.PreserveFailure(ctx, op)
}
func (s *Service) PreserveSnapshotFailure(ctx context.Context, op func() error) error {
	return runtime.PreserveSnapshotFailure(ctx, op)
}
func (s *Service) PreserveHealthFailure(ctx context.Context, op func() error) error {
	return runtime.PreserveHealthFailure(ctx, op)
}
func (s *Service) NewRuntimeState() *runtime.State { return runtime.NewState() }
