package piscine

import (
	"context"
	"time"
)

func FetchWithTimeout(ctx context.Context, workDuration time.Duration) (string, error) {
	select {
	case <-time.After(workDuration):
		return "done", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func ProcessAll(ctx context.Context, items []string, fn func(string)) error {
	for _, item := range items {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		fn(item)
	}
	return nil
}
