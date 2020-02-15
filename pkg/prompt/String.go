// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package prompt

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func String(question string, stdout bool, loop bool) (string, error) {
	if len(question) > 0 {
		if stdout {
			_, _ = fmt.Fprintf(os.Stdout, "%s: ", question)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "%s: ", question)
		}
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
			// add new line
			if stdout {
				_, _ = fmt.Fprintln(os.Stdout, "")
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "")
			}
			// stop waiting
			done <- true
		}
	}()

	go func() {
		for {
			str, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				inputErr = fmt.Errorf("error reading string from terminal: %w", err)
				break
			}
			str = strings.TrimSpace(str)
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

	return value, inputErr
}
