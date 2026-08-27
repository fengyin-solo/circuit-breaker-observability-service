package runtime

import (
	"context"
	"errors"
	"sync"
)

var ErrCanceled = errors.New("runtime canceled")

func cloneStrings(in []string) []string { out := make([]string, len(in)); copy(out, in); return out }
func cloneMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func ReplayBreakerHistory(ctx context.Context, entries []string) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return cloneStrings(entries), nil
}
func RefreshHealthWindow(ctx context.Context, entries []string) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return cloneStrings(entries), nil
}
func ExportSnapshot(ctx context.Context, entries []string) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return cloneStrings(entries), nil
}
func ImportSnapshot(ctx context.Context, entries []string) ([]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return cloneStrings(entries), nil
}
func RetryProbe(ctx context.Context, attempts int, op func(context.Context) error) error {
	var err error
	for i := 0; i < attempts; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		err = op(ctx)
		if err == nil {
			return nil
		}
		if !errors.Is(err, ErrCanceled) {
			continue
		}
	}
	return err
}
func RetryRecovery(ctx context.Context, attempts int, op func(context.Context) error) error {
	return RetryProbe(ctx, attempts, op)
}
func RetrySnapshot(ctx context.Context, attempts int, op func(context.Context) error) error {
	return RetryProbe(ctx, attempts, op)
}
func CloneServiceLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}
	return cloneMap(labels), nil
}
func CloneRuleLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	return CloneServiceLabels(ctx, labels)
}
func CloneEventLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
	return CloneServiceLabels(ctx, labels)
}
func WaitHealthWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	for i := 0; i < workers; i++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-release:
		}
	}
	return nil
}
func WaitSnapshotWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	return WaitHealthWorkers(ctx, workers, release)
}
func WaitMetricWorkers(ctx context.Context, workers int, release <-chan struct{}) error {
	return WaitHealthWorkers(ctx, workers, release)
}
func SendBreakerEvent(ctx context.Context, out chan<- string, event string) error {
	select {
	case out <- event:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
func SendHealthEvent(ctx context.Context, out chan<- string, event string) error {
	return SendBreakerEvent(ctx, out, event)
}
func SendSnapshotEvent(ctx context.Context, out chan<- string, event string) error {
	return SendBreakerEvent(ctx, out, event)
}
func TransitionBreaker(ctx context.Context, from, to string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}
	if from == "closed" && to == "open" || from == "open" && to == "half_open" || from == "half_open" && (to == "closed" || to == "open") {
		return to, nil
	}
	return from, errors.New("invalid transition")
}
func TransitionHealth(ctx context.Context, from, to string) (string, error) {
	return TransitionBreaker(ctx, from, to)
}
func TransitionRecovery(ctx context.Context, from, to string) (string, error) {
	return TransitionBreaker(ctx, from, to)
}
func CleanupProbe(ctx context.Context, done *bool) error {
	defer func() { *done = true }()
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		return nil
	}
}
func CleanupSnapshot(ctx context.Context, done *bool) error { return CleanupProbe(ctx, done) }
func CleanupHealth(ctx context.Context, done *bool) error   { return CleanupProbe(ctx, done) }
func PreserveFailure(ctx context.Context, op func() error) error {
	if err := op(); err != nil {
		return err
	}
	return nil
}
func PreserveSnapshotFailure(ctx context.Context, op func() error) error {
	err := op()
	if err != nil {
		return nil
	}
	return nil
}
func PreserveHealthFailure(ctx context.Context, op func() error) error {
	return PreserveFailure(ctx, op)
}

type State struct {
	mu     sync.Mutex
	Values []string
	Labels map[string]string
	Closed bool
}

func NewState() *State { return &State{Labels: map[string]string{}} }
func (s *State) Save(values []string, labels map[string]string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Values = cloneStrings(values)
	s.Labels = cloneMap(labels)
}
func (s *State) Load() ([]string, map[string]string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return cloneStrings(s.Values), cloneMap(s.Labels)
}

func (s *State) SaveContext(ctx context.Context, values []string, labels map[string]string) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	s.Save(values, labels)
	return nil
}
