package executil

import (
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

func TestCloneCmd(t *testing.T) {
	cmd, err := ParseCmd(os.Args[0] + " -test.run=Test_SimpleSubprocess -- Howdy")

	if err != nil {
		t.Fatalf("ParseCmd() returned err: +%v", err)
	}

	cmd.Env = append(cmd.Env, "GO_RUNNING_SUBPROCESS=1")

	clonedCmd := CloneCmd(cmd)

	if cmd == clonedCmd {
		t.Fatalf("expected two different commands, but got the same one")
	}

	out, err := clonedCmd.Output()
	if err != nil {
		t.Fatalf("got err running cloned command: %+v", err)
	}

	want := "Howdy there"
	got := string(out)

	if want != got {
		t.Errorf("cmd output = %q; want %q", got, want)
	}
}

// -----------------------------------------------------------------------------
// Sample programs to use for cmd testing

func Test_SimpleSubprocess(t *testing.T) {
	if os.Getenv("GO_RUNNING_SUBPROCESS") != "1" {
		return
	}

	greeting := os.Args[len(os.Args)-1]
	fmt.Printf("%s there", greeting)
	os.Exit(0)
}
