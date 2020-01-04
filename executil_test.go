package executil

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func TestParseCmd(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		cmd, err := ParseCmd("")

		if err != ErrEmptyCommandString {
			t.Errorf("Expected error for empty command string, but got: %+v", err)
		}

		if cmd != nil {
			t.Errorf("Expected cmd to be nil but got: %+v", cmd)
		}
	})

	t.Run("without args", func(t *testing.T) {
		cmd, err := ParseCmd(os.Args[0])

		if err != nil {
			t.Errorf("ParseCmd() returned err: %+v", err)
		}

		if cmd == nil {
			t.Fatalf("ParseCmd() returned nil cmd")
		}

		want := os.Args[0]
		got := cmd.Path

		if want != got {
			t.Errorf("cmd.Path = %s; want %s", got, want)
		}

		if 1 != len(cmd.Args) {
			t.Errorf("expected cmd.Args to be empty, but got: %v", cmd.Args)
		}
	})

	t.Run("with args", func(t *testing.T) {
		cmd, err := ParseCmd(os.Args[0] + " -some=value")

		if err != nil {
			t.Errorf("ParseCmd() returned err: %+v", err)
		}

		if cmd == nil {
			t.Fatalf("ParseCmd() returned nil cmd")
		}

		want := os.Args[0]
		got := cmd.Path

		if want != got {
			t.Errorf("cmd.Path = %s; want %s", got, want)
		}

		if 2 != len(cmd.Args) {
			t.Errorf("expected 2 parts, but got %d", len(cmd.Args))
		}
	})
}

// -----------------------------------------------------------------------------
// Sample programs to use for cmd testing

func Test_SimpleSubprocess(t *testing.T) {
	if os.Getenv("GO_RUNNING_SUBPROCESS") != "1" {
		return
	}

	var greeting string
	flag.StringVar(&greeting, "greeting", "Hello", "greeting")

	fmt.Printf("%s there", greeting)
}
