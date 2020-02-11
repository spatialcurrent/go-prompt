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
	"sync"
	"syscall"
)

func String(question string, stdout bool) (string, error) {
	if len(question) > 0 {
		if stdout {
			fmt.Fprintf(os.Stdout, "%s: ", question)
		} else {
			fmt.Fprintf(os.Stderr, "%s: ", question)
		}
	}

	value := ""
	var inputErr error

	wg := &sync.WaitGroup{}
	wg.Add(1)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGPIPE)
	go func() {
		s := <-signals
		signal.Stop(signals)
		inputErr = fmt.Errorf("received signal %q", s)
		// add new line
		if stdout {
			fmt.Fprintln(os.Stdout, "")
		} else {
			fmt.Fprintln(os.Stderr, "")
		}
		// stop waiting
		wg.Done()
	}()

	go func() {
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			inputErr = fmt.Errorf("error reading string from terminal: %w", err)
		} else {
			value = strings.TrimSpace(str)
		}
		// stop waiting
		wg.Done()
	}()

	wg.Wait()

	return value, inputErr
}
