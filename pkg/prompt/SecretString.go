// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package prompt

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func SecretString(question string, stdout bool, loop bool) (string, error) {
	if len(question) > 0 {
		if stdout {
			_, _ = fmt.Fprintf(os.Stdout, "%s: ", question)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "%s: ", question)
		}
	}

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return "", fmt.Errorf("error making terminal raw: %w", err)
	}

	value := ""
	var inputErr error

	done := make(chan bool, 1)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	go func() {
		for s := range signals {
			signal.Stop(signals)
			inputErr = fmt.Errorf("received signal %q", s)
			done <- true
		}
	}()

	go func() {
		for {
			b, err := terminal.ReadPassword(int(os.Stdin.Fd()))
			if err != nil {
				inputErr = fmt.Errorf("error reading secret string from terminal: %w", err)
				break
			}
			str := string(bytes.TrimSpace(b))
			if len(str) == 0 && loop {
				continue
			}
			value = str
			break
		}
		// stop waiting
		done <- true
		signal.Stop(signals)
		close(signals)
	}()

	<-done

	signal.Stop(signals)

	if oldState != nil {
		_ = terminal.Restore(0, oldState)
	}

	// add new line
	if stdout {
		_, _ = fmt.Fprintln(os.Stdout, "")
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "")
	}

	return value, inputErr
}
