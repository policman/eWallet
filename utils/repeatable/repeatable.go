package repeatable

import (
	"log/slog"
	"time"
)

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--
			continue
		}
		return nil
	}

	slog.Error("error with connect to postgresql in 5 times", err)
	return err

}
