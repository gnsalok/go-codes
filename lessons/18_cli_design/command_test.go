package clidesign

import (
	"bytes"
	"testing"
)

func TestRunWritesGreeting(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run([]string{"-name", "Gopher"}, &stdout, &stderr)

	if code != ExitOK || stdout.String() != "hello, Gopher\n" || stderr.Len() != 0 {
		t.Fatalf("Run() code=%d stdout=%q stderr=%q", code, stdout.String(), stderr.String())
	}
}

func TestRunReportsUsageError(t *testing.T) {
	var stdout, stderr bytes.Buffer
	code := Run(nil, &stdout, &stderr)

	if code != ExitUsage || stderr.String() == "" {
		t.Fatalf("Run() code=%d stderr=%q; want usage error", code, stderr.String())
	}
}
