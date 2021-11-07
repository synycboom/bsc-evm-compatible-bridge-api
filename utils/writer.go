package utils

import (
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
)

const maxRetries = 10
const timeout = time.Second * 10

func TryWrite(w io.Writer, bz []byte) error {
	var err error
	cont := true

	for retry := 0; cont; retry++ {
		if retry > maxRetries {
			return fmt.Errorf("reached max retries limit. Last error was: %v", err)
		}

		done := make(chan bool)
		go func() {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("write call panicked: %v", r)
				}
				done <- true
			}()

			n := 0
			for n < len(bz) {
				i, err := w.Write(bz)
				if err != nil {
					return
				}
				n += i
			}
		}()

		select {
		case <-done:
			if err == nil {
				cont = false
			}
		case <-time.After(timeout):
			return fmt.Errorf("reached timeout (10s)")
		}
	}

	if err != nil {
		return errors.Wrap(err, "TryWrite error during response write")
	}
	return nil
}
