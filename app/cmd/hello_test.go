package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// TestHelloCmd verifies that the hello subcommand prints the expected greeting.
func TestHelloCmd(t *testing.T) {
	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	helloCmd.Run(helloCmd, []string{})

	if err := w.Close(); err != nil {
		t.Fatalf("failed to close pipe writer: %v", err)
	}
	os.Stdout = origStdout

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("failed to read captured stdout: %v", err)
	}

	got := strings.TrimSpace(buf.String())
	want := "Hello, World!!!"
	if got != want {
		t.Errorf("hello command output = %q, want %q", got, want)
	}
}

// TestHelloCmdRegistered ensures the hello command is wired up to the root command.
func TestHelloCmdRegistered(t *testing.T) {
	for _, c := range rootCmd.Commands() {
		if c.Name() == "hello" {
			return
		}
	}
	t.Fatal("expected 'hello' subcommand to be registered on rootCmd")
}
